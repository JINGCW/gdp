package demapstructure

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type DecodeHookFunc interface{}

//DocodeHookFuncType ia a DecodeHookFunc which has completely information about
//the source and target types.
type DecodeHookFuncType func(reflect.Type, reflect.Type, interface{}) (interface{}, error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
// source and target types.
type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)

type DecoderConfig struct {
	//DecodeHook, if set, will be called before any decoding and any
	//type conversion (if WeaklyTypedInput is on). This lets you modify
	//the values before they're set down onto the resulting struct.
	//
	//If an error is returned, the entire decode will fail with that
	//error.
	DecodeHook DecodeHookFunc

	//If ErrorUnused is true, then it is an error for there to exist
	//keys in the original map that were unused in the decoding process
	//(extra keys).
	ErrorUnused bool

	//ZeroFields, if set to true, will zero fields before writing them.
	//For example, a map will be emptied before decoded values are put in
	//it. If this is false, a map will be merged.
	ZeroFields bool

	//If WeaklyTypedInput is true, the decoder will make the following
	//"weak" conversion:
	//
	//   - bools to string (true = "1", false = "0")
	//   - numbers to string (base 10)
	//   - bools to int/uint (true = 1, false = 0)
	//   - strings to int/uint (base implied by prefix)
	//   - int to bool (true if value != 0)
	//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	//     FALSE, false, False. Anything else is an error)
	//   - empty array = empty map and vice versa
	//   - negative numbers to overflowed uint values (base 10)
	//   - slice of maps to a merged map
	//   - single values are converted to slices if required. Each
	//     element is weakly decoded. For example: "4" can become []int{4}
	//     if the target type is an int slice.
	//
	WeaklyTypedInput bool

	//Metadata is the struct that will contain extra metadata about
	//the decoding. If this is nil, then no metadata will be tracked.
	Metadata *Metadata

	//Result is a pointer to the struct that will contain the decoded
	//value.
	Result interface{}

	// The tag name that mapstructure reads for field names. This
	// defaults to "mapstructure"
	TagName string
}

type Metadata struct {
	//Keys are the keys of the structure which were successfully decoded
	Keys []string

	//Unused is a slice of keys that were found in the raw value but
	//weren't decoded since there was no matching field in the result interface
	Unused []string
}

//A Decoder takes a raw interface value and turns it into structured
//data, keeping track of rich error information along the way in case
//anything goes wrong. Unlike the basic top-level Decode method, you can
//more finely control how the Decoder behaves using the DecoderConfig
//structure. The top-level Decode method is just a convenience that sets
//up the most basic Decoder.
type Decoder struct {
	config *DecoderConfig
}

//Decode takes an input structure and uses reflection to translate it to
//the output structure. output must be a pointer to a map or struct.
func Decode(input interface{}, output interface{}) error {
	config := &DecoderConfig{
		Metadata: nil,
		Result:   output,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(config *DecoderConfig) (*Decoder, error) {
	val := reflect.ValueOf(config.Result)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}

	val = val.Elem()
	if !val.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if config.Metadata != nil {
		if config.Metadata.Keys == nil {
			config.Metadata.Keys = make([]string, 0)
		}
		if config.Metadata.Unused == nil {
			config.Metadata.Unused = make([]string, 0)
		}
	}

	if config.TagName == "" {
		config.TagName = "mapstructure"
	}
	result := &Decoder{config: config}
	return result, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (d *Decoder) Decode(input interface{}) error {
	return d.decode("", input, reflect.ValueOf(d.config.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (d *Decoder) decode(name string, input interface{}, outVal reflect.Value) error {
	var inputVal reflect.Value
	if input != nil {
		inputVal = reflect.ValueOf(input)
		if inputVal.Kind() == reflect.Ptr && inputVal.IsNil() {
			input = nil
		}
	}

	if input == nil {
		if d.config.ZeroFields {
			outVal.Set(reflect.Zero(outVal.Type()))
			if d.config.Metadata != nil && name != "" {
				d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
			}
		}
		return nil
	}

	if !inputVal.IsValid() {
		outVal.Set(reflect.Zero(outVal.Type()))
		if d.config.Metadata != nil && name != "" {
			d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
			return nil
		}
	}

	//var err error
	//outputKind:=getKind(outVal)
	//switch outputKind {
	//
	//}

	return nil
}

func (d *Decoder) decodeBool(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataKind := dataVal.Kind()

	switch {
	case dataKind == reflect.Bool:
		val.SetBool(dataVal.Bool())
	case dataKind == reflect.Int && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Int() != 0)
	case dataKind == reflect.Uint && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Uint() != 0)
	case dataKind == reflect.Float32 && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Float() != 0)
	case dataKind == reflect.String && d.config.WeaklyTypedInput:
		b, err := strconv.ParseBool(dataVal.String())
		if err == nil {
			val.SetBool(b)
		} else if dataVal.String() == "" {
			val.SetBool(false)
		} else {
			fmt.Errorf("'%s' expected type '%s', got unconvertible type '%s'",
				name, val.Type(), dataVal.Type())
		}
	}
	return nil
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (d *Decoder) decodeBasic(name string, data interface{}, val reflect.Value) error {
	if val.IsValid() && val.Elem().IsValid() {
		return d.decode(name, data, val.Elem())
	}

	dataVal := reflect.ValueOf(data)
	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if dataVal.Kind() == reflect.Ptr && dataVal.Type().Elem() == val.Type() {
		dataVal = reflect.Indirect(dataVal)
	}

	if !dataVal.IsValid() {
		dataVal = reflect.Zero(val.Type())
	}

	dataValType := dataVal.Type()
	if !dataValType.AssignableTo(val.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'",
			name, val.Type(), dataValType)
	}

	val.Set(dataVal)
	return nil
}

func (d *Decoder) decodeString(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataKind := getKind(dataVal)

	converted := true
	switch {
	case dataKind == reflect.String:
		val.SetString(dataVal.String())
	case dataKind == reflect.Bool && d.config.WeaklyTypedInput:
		if dataVal.Bool() {
			val.SetString("1")
		} else {
			val.SetString("0")
		}
	case dataKind == reflect.Int && d.config.WeaklyTypedInput:
		val.SetString(strconv.FormatInt(dataVal.Int(), 10))
	case dataKind == reflect.Float32 && d.config.WeaklyTypedInput:
		val.SetString(strconv.FormatFloat(dataVal.Float(), 'f', -1, 64))
	case dataKind == reflect.Slice && d.config.WeaklyTypedInput,
		dataKind == reflect.Array && d.config.WeaklyTypedInput:
		dataType := dataVal.Type()
		elemKind := dataType.Elem().Kind()
		switch elemKind {
		case reflect.Uint8:
			var uints []uint8
			if dataKind == reflect.Array {
				uints = make([]uint8, dataVal.Len(), dataVal.Len())
				for i := range uints {
					uints[i] = dataVal.Index(i).Interface().(uint8)
				}
			} else {
				uints = dataVal.Interface().([]uint8)
			}
			val.SetString(string(uints))
		default:
			converted = false
		}
	default:
		converted = false
	}
	if !converted {
		return fmt.Errorf("'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}
	return nil
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

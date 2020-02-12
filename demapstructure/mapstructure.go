package demapstructure

import "reflect"

type DecodeHookFunc interface {}

//DocodeHookFuncType ia a DecodeHookFunc which has completely information about
//the source and target types.
type DecodeHookFuncType func(reflect.Type,reflect.Type,interface{})(interface{},error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
// source and target types.
type DecodeHookFuncKind func(reflect.Kind,reflect.Kind,interface{})(interface{},error)

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

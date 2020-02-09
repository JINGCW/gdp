package dereflect

import (
	"strconv"
	"unsafe"
)
// Type is the representation of a Go type.
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.

//Type values are comparable, such as with the == operator,
//so they can be used as map keys.
//Two Type values are equal if they represent identical types.
type Type interface {
	//Methods applicable to all types.

	//Align returns the alignment in bytes of a value of
	//this type when allocated in memory.
	Align()int

	//FieldAlign returns the alignment in bytes of a value of
	//this type when used as a field in a struct.
	FieldAlign()int

	//Method returns the i'th method in the type's method set.
	//It panics if i is not in the range [0,NumMethod()].
	//
	//For a non-interface type T or *T, the returned Method's Type and Func
	//fields describe a function whose first argument is the receiver.
	//
	//For an interface type, the returned Method's Type field gives the
	//method signature, without a receiver, and the Func field is nil.
	//
	//Only exported methods are accessible and they are sorted in
	//lexicographic order.
	Method(int)Method

	//MethodByName returns the method with that name in the type's
	//method set and boolean indicating if the method was found.
	//
	//For a non-interface type T or *T, the returned Method's Type and Func
	//fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	MethodByName(string)(Method,string)

	//returns the number of exported methods in the type's method set.
	NumMethod()int

	//returns the type's name within its package for a defined type.
	//For other (non-defined) types it returns the empty string.
	Name()string

	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
	PkgPath() string

	//returns the number of bytes needed to store
	//a value of the given type; it is analogous to unsafe.Sizeof.
	Size()uintptr

	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
	String() string

	//returns the specific kind of this type.
	Kind()Kind

	//reports whether the type implements the interface type u.
	Implements(u Type)bool

	//reports whether a value of the type is assignable to type u.
	AssignableTo(u Type)bool

	//reports whether a value of the type is convertible to type u.
	ConvertibleTo(u Type)bool

	//reports whether values of this type are comparable.
	Comparable()bool

	// Methods applicable only to some types, depending on Kind.
	// The methods allowed for each kind are:
	//
	//	Int*, Uint*, Float*, Complex*: Bits
	//	Array: Elem, Len
	//	Chan: ChanDir, Elem
	//	Func: In, NumIn, Out, NumOut, IsVariadic.
	//	Map: Key, Elem
	//	Ptr: Elem
	//	Slice: Elem
	//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField

	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	Bits()int

	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	ChanDir()ChanDir

	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic()bool

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem()Type

	// Field returns a struct type's i'th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int)structField

	// FieldByIndex returns the nested field corresponding
	// to the index sequence. It is equivalent to calling Field
	// successively for each index i.
	// It panics if the type's Kind is not Struct.
	FieldByIndex(index []int)structField

	// FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	FieldByName(name string)(structField,bool)

	// FieldByNameFunc returns the struct field with a name
	// that satisfies the match function and a boolean indicating if
	// the field was found.
	//
	// FieldByNameFunc considers the fields in the struct itself
	// and then the fields in any embedded structs, in breadth first order,
	// stopping at the shallowest nesting depth containing one or more
	// fields satisfying the match function. If multiple fields at that depth
	// satisfy the match function, they cancel each other
	// and FieldByNameFunc returns no match.
	// This behavior mirrors Go's handling of name lookup in
	// structs containing embedded fields.
	FieldByNameFunc(match func(string2 string)bool)(structField,bool)

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int)Type

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key()Type

	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len()int

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField()int

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn()int

	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut()int

	// Out returns the type of a function type's i'th output parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(i int)Type

	common()*rtype
	uncommon()*uncommonType
}

// uncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
type uncommonType struct {
	pkgPath nameOff // import path; empty for built-in types like int, string
	mcount  uint16  // number of methods
	xcount  uint16  // number of exported methods
	moff    uint32  // offset from this uncommontype to [mcount]method
	_       uint32  // unused
}


//Method represents a single method.
type Method struct {
	Name string
	PkgName string

	Type Type
	Func Value//func with receiver as first argument
	Index int//index for Type.Method
}

// ChanDir represents a channel type's direction.
type ChanDir int

const (
	RecvDir ChanDir             = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)

//arrayType represents a fixed array type.
type arrayType struct {
	rtype
	elem *rtype//array element type
	slice *rtype//slice type
	len uintptr
}

type chanType struct {
	rtype
	elem *rtype//channel element type
	dir uintptr//channel direction (ChanDir)
}

//type StructField struct {
//	Name string
//	// PkgPath is the package path that qualifies a lower case (unexported)
//	// field name. It is empty for upper case (exported) field names.
//	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
//	PkgPath string
//
//	Type Type
//	Tag StructTag
//	Offset uintptr //offset within struct, in bytes
//	Index []int //index sequence for Type.FieldByIndex
//	Anonymous bool //is an embedded field
//}

type StructTag string

type tflag uint8

const (
	//tflagUncommon means that there is a pointer, *uncommonType,
	//just beyond the outer type structure.
	//
	//For example, if t.Kind()==Struct and t.tflag&tflagUncommon!=0,
	//then t has uncommonType data and it can be accessed as:
	//
	//type tUncommon struct{
	//    structType
	//    u uncommonType
	//}
	//u:=&(*tUncommon)(unsafe.Pointer(t)).u
	tflagUncommon tflag=1<<0

	//tflagExtraStar means the name in the str field has an
	//extraneous '*' prefix. This is because for most types T in
	//a program, the type *T also exists and reusing the str data
	//saves binary size.
	tflagExtraStar tflag=1<<1

	//tflagNamed means the type has a name.
	tflagNamed tflag=1<<2
)

type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type textOff int32 // offset from top of text section

type structTypeUncommon struct {
	structType
	u uncommonType
}

////ptrType represents a pointer type.
//type ptrType struct {
//	rtype
//	elem *rtype//pointer element (pointed at) type
//}

//Struct field
type structField struct {
	name name//name is always non-empty
	typ *rtype//type of field
	offsetEmbed uintptr//byte offset of field<<1|isEmbedded
}

func (f*structField)offset()uintptr{
	return f.offsetEmbed>>1
}

func (f*structField)embedded()bool{
	return f.offsetEmbed&1!=0
}

//structType represents a struct type.
type structType struct {
	rtype
	pkgPath name
	fields []structField//sorted by offset
}

//sliceType represents a slice type.
type sliceType struct {
	rtype
	elem *rtype//slice element type
}

//funcType represents a function type.
//
//A *rtype for each in and out parameter is stored in an array that
//directly follows the funcType (and possibly its uncommonType). So
//a function type with one method, one input, and one output is:
//
//struct {
//    funcType
//    uncommonType
//    [2]*rtype //[0] is in, [1] is out
//}
type funcType struct {
	rtype
	inCount uint16
	outCount uint16
}

//imethod represents a method on an interface type
type imethod struct {
	name nameOff//name of method
	typ typeOff//.(*FuncType) underneath
}

//interfaceType represents an interface type.
type interfaceType struct {
	rtype
	pkgPath name//import path
	methods []imethod
}

//mapType represents a map type.
type mapType struct {
	rtype
	key *rtype//map key type
	elem *rtype//map element (value) type
	bucket *rtype//internal bucket structure
	keysize uint8//size of key slot
	valuesize uint8//size of value slot
	bucketsize uint16//size of bucket
	flags uint32
}

//ptrType represents a pointer type.
type ptrType struct {
	rtype
	elem *rtype//pointer element (pointed at) type
}

func (t*rtype)uncommon()*uncommonType{
	if t.tflag&tflagUncommon==0{
		return nil
	}
	switch t.Kind() {
	case Struct:
		return &(*structTypeUncommon)(unsafe.Pointer(t)).u
	case Ptr:
		type u struct {
		ptrType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Func:
		type u struct {
		funcType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Slice:
		type u struct {
		sliceType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Array:
		type u struct {
		arrayType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Chan:
		type u struct {
		chanType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Map:
		type u struct {
		mapType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case Interface:
		type u struct {
		interfaceType
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	default:
		type u struct {
		rtype
		u uncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	}
}

func (t*rtype)String()string{
	return "nil"//todo
}

func (t*rtype)Size()uintptr{return t.size}

func (t*rtype)Bits()int{
	if t==nil{
		panic("derefect: Bits of nil Type")
	}
	k:=t.Kind()
	if k<Int||k>Complex128{
		panic("dereflect: Bits of non-arithmetic Type"+t.String())
	}
	return int(t.size)*8
}

func (t*rtype)Align()int{return int(t.align)}

func (t*rtype)FieldAlign()int  {
	return int(t.fieldAlign)
}

//func (t *rtype) Kind() exists

func (t*rtype)pointers()bool  {
	return t.ptrdata!=0
}

func (t*rtype)common()*rtype{return t}

//Method on non-interface type
type method struct {
	name nameOff//name of method
	mtyp typeOff//method type (without receiver)
	ifn textOff//fn used in interface call (one-word receiver)
	tfn textOff//fn used for normal method call
}

func (t*rtype)exportedMethods()[]method{
	ut:=t.uncommon()
	if ut==nil{
		return nil
	}
	return ut.exportedMethods()
}

type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
type rtype struct {
	size       uintptr
	ptrdata    uintptr  // number of bytes in the type that can contain pointers
	hash       uint32   // hash of type; avoids computation in hash tables
	tflag      tflag    // extra type information flags
	align      uint8    // alignment of variable with this type
	fieldAlign uint8    // alignment of struct field with this type
	kind       uint8    // enumeration for C
	alg        *typeAlg // algorithm table
	gcdata     *byte    // garbage collection data
	str        nameOff  // string form
	ptrToThis  typeOff  // type for pointer to this type, may be zero
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

const (
	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6 // Type.gc points to GC program
	kindMask        = (1 << 5) - 1
)

//String returns the name of k.
func(k Kind)String()string{
	if int(k)<len(kindNames){
		return kindNames[k]
	}
	return "kind"+strconv.Itoa(int(k))
}

var kindNames=[]string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Ptr:           "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}

func(t*uncommonType)methods()[]method{
	if t.mcount==0{
		return nil
	}
	return (*[1<<16]method)(add(unsafe.Pointer(t),uintptr(t.moff),"t.mcount>0"))[:t.mcount:t.mcount]
}

func(t*uncommonType)exportedMethods()[]method{
	if t.xcount==0{
		return nil
	}
	return (*[1<<16]method)(add(unsafe.Pointer(t),uintptr(t.moff),"t.mcount>0"))[:t.xcount:t.xcount]
}

func (t *rtype) Kind() Kind { return Kind(t.kind & kindMask) }

// ifaceIndir reports whether t is stored indirectly in an interface value.
func ifaceIndir(t *rtype) bool {
	return t.kind&kindDirectIface == 0
}

//name is an encoded type name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the name is exported
//	1<<1 tag data follows the name
//	1<<2 pkgPath nameOff follows the name and tag
//
// The next two bytes are the data length:
//
//	 l := uint16(data[1])<<8 | uint16(data[2])
//
// Bytes [3:3+l] are the string data.
//
// If tag data follows then bytes 3+l and 3+l+1 are the tag length,
// with the data following.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
type name struct {
	bytes *byte
}

func (n name) data(off int, whySafe string) *byte {
	return (*byte)(add(unsafe.Pointer(n.bytes), uintptr(off), whySafe))
}

func (n name) isExported() bool {
	return (*n.bytes)&(1<<0) != 0
}

func (n name) nameLen() int {
	return int(uint16(*n.data(1, "name len field"))<<8 | uint16(*n.data(2, "name len field")))
}

func (n name) tagLen() int {
	if *n.data(1, "name flag field")&(1<<1) == 0 {
		return 0
	}
	off := 3 + n.nameLen()
	return int(uint16(*n.data(off, "name taglen field"))<<8 |
		uint16(*n.data(off+1, "name taglen field")))
}

func (n name) name() (s string) {
	if n.bytes == nil {
		return
	}
	b := (*[4]byte)(unsafe.Pointer(n.bytes))
	hdr := (*stringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(&b[3])
	hdr.Len = int(b[1])<<8 | int(b[2])
	return s
}

func (n name) tag() (s string) {
	tl := n.tagLen()
	if tl == 0 {
		return ""
	}
	nl := n.nameLen()
	hdr := (*stringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(n.data(3+nl+2, "non empty-string"))
	hdr.Len = tl
	return s
}

func (n name) pkgPath() string {
	if n.bytes == nil || *n.data(0, "name flag field")&(1<<2) == 0 {
		return ""
	}
	off := 3 + n.nameLen()
	if tl := n.tagLen(); tl > 0 {
		off += tl
	}
	var nameOff int32
	copy((*[4]byte)(unsafe.Pointer(&nameOff))[:],
		(*[4]byte)(unsafe.Pointer(n.data(off, "name offset field")))[:])
	//pkgPathName:=name{(*byte)(unsafe.Pointer())}
	return "cannot use resolveTypeOff of reflect"
}

func Rount(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

func NewName(n, tag string, exported bool) name {
	if len(n) > 1<<16-1 {
		panic("reflect.nameFrom: name too long" + n)
	}
	if len(tag) > 1<<16-1 {
		panic("reflect.tagFrom: tag too long" + tag)
	}
	var bits byte
	l := 1 + 2 + len(n)
	if exported {
		bits |= 1 << 0
	}
	if len(tag) > 0 {
		l += 2 + len(tag)
		bits |= 1 << 1
	}
	b := make([]byte, l)
	b[0] = bits
	b[1] = uint8(len(n) >> 8)
	b[2] = uint8(len(n))
	copy(b[3:], n)
	if len(tag) > 0 {
		tb := b[3+len(n):]
		tb[0] = uint8(len(tag) >> 8)
		tb[1] = uint8(len(tag))
		copy(tb[2:], tag)
	}
	return name{bytes: &b[0]}
}

/*
 The compiler knows the exact layout of all the data structures above.
 The compiler dose not know the data structures and methods below.
 */

//Struct field describes a single field in a struct.
type StructField struct {
	Name string
	PkgPath string

	Type Type
	Tag StructTag
	Offset uintptr
	Index []int
	Anonymous bool
}

func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

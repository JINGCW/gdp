package main

import (
	"fmt"
	"reflect"
)

//
//const (
//	mask uintptr = 1<<5 - 1
//)
//
//type Kind uint
//
//const (
//	Invalid Kind = iota
//	Bool
//	Int
//)
//
//const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
//const temp  = (^uintptr(0) >> 1)
//const MaxUintptr = ^uintptr(0)

type rwer interface {
	read()string
	write()string
}

type reader interface {
	Read() string
}

type concreate_type struct {
	Val string
}

func (e *concreate_type)Read()string{
	return e.Val
}

type test interface {
	name()string
}



func main() {
	input := map[string]interface{}{
		"name":   123,                      // number => string
		"age":    "42",                     // string => number
		"emails": map[string]interface{}{}, // empty map => empty array
	}
	//var v *int
	vv:=reflect.ValueOf(input["emails"]).Type().Elem().Kind()
	fmt.Println(vv)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	//input := map[string]interface{}{
	//	"name":   "Mitchell",
	//	"age":    91,
	//	"emails": []string{"one", "two", "three"},
	//	"extra": map[string]string{
	//		"twitter": "mitchellh",
	//	},
	//}

	//var result Person=Person{
	//
	//}
	//dataValKind:=dataValType.Key().Kind()

	//var r1 reader
	//r1=&concreate_type{"haha"}
	//fmt.Println(reflect.ValueOf(iv).Elem())
	//var iv type_interface=&concreate_type{"somthing"}
	//var eface interface{}= (interface{})(*(*interface {
	//	M()string
	//})(unsafe.Pointer(reflect.ValueOf(err).Pointer())))
	//var r2 rwer=r1
	//out:="abcd"
	//out2:=dereflect.ValueOf(out)
	//fmt.Println(unsafe.Sizeof(out))
	//fmt.Println(unsafe.Sizeof(out2))
	//fmt.Printf("%x", temp)
	//fmt.Println(HashStr("abc"))
}

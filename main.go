package main

import (
	"hash_collections/dereflect"
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

type type_interface interface {
	type_func() string
}

type concreate_type struct {
	val string
}

func (e *concreate_type)type_func()string{
	return e.val
}

func main() {
	var out  =concreate_type{"concreate_value"}
	//out=&concreate_type{"concreate_value"}
	//println(reflect.ValueOf(out).Type())
	println(dereflect.ValueOf(out).Kind())
	//out:="abcd"
	//out2:=dereflect.ValueOf(out)
	//fmt.Println(unsafe.Sizeof(out))
	//fmt.Println(unsafe.Sizeof(out2))
	//fmt.Printf("%x", temp)
	//fmt.Println(HashStr("abc"))
}

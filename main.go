package main

import "hash_collections/dereflect"

const (
	mask uintptr = 1<<5 - 1
)

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
)

const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
const temp  = (^uintptr(0) >> 1)
const MaxUintptr = ^uintptr(0)

func main() {
	println(dereflect.RecvDir)
	println(dereflect.SendDir)
	println(dereflect.BothDir)
	//out:="abcd"
	//out2:=dereflect.ValueOf(out)
	//fmt.Println(unsafe.Sizeof(out))
	//fmt.Println(unsafe.Sizeof(out2))
	//fmt.Printf("%x", temp)
	//fmt.Println(HashStr("abc"))
}

package main

import "hash_collections/src/asm_x86/g"

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

//}

func main() {
	g1,g2:=g.G_innerouter()
	println(g1,g2)
	//fmt.Println(add_sub_mul.Output2(3,4))
	//gid.GoID()
	//var m = map[int]int{}
	//m[43] = 1
	//var n = map[string]int{}
	//n["abc"] = 1
	//println(m, n)
	//fmt.Println(add_sub_mul.ArraySum([]int64{12, 3, 4}))
	//deracy.SockServer()
	//deracy.SockClient()
	//fmt.Printf("%x\n",-1&-1)
	//desync.MuDemo()
	//config.ReadInConfig()
	//fmt.Println(viper.GetStringSlice("cluster_hp"))
	//_=dego.Cclient(viper.GetStringSlice("cluster_hp"), viper.GetString("passwd"))
	//if e!=nil{
	//	panic(e)
	//}
	//s:=dego_redis.Usr{"hello","123@qq.com",123}
	//bytes,err:=json.Marshal(s)
	//if err!=nil{
	//	panic(err)
	//}
	//out:=dego_redis.Usr{}
	//_=json.Unmarshal(bytes,&out)
	//fmt.Printf("%v",reflect.Indirect(reflect.ValueOf(out)).Kind())
	//fmt.Printf("%v",out)

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

package main
//
//import (
//	"fmt"
//	"os"
//	"path"
//	"plugin"
//	"syscall"
//)
//
//func main1() {
//	dll := syscall.NewLazyDLL("dll_c.dll")
//	demo := dll.NewProc("Demo")
//	out, err, msg := demo.Call()
//	fmt.Println(out)
//	fmt.Println(err)
//	fmt.Println(msg)
//}
//
//func main() {
//	f, err := os.OpenFile(".", os.O_RDONLY, 0666)
//	if err != nil {
//		panic(err)
//	}
//	fi, err := f.Readdir(-1)
//	if err != nil {
//		panic(err)
//	}
//	out := make([]os.FileInfo, 0)
//
//	for _, ff := range fi {
//		println(ff.Name())
//		if ff.IsDir() || path.Ext(ff.Name()) != ".a" {
//			continue
//		}
//		out = append(out, ff)
//		pdll, err := plugin.Open(ff.Name())
//		if err != nil {
//			fmt.Println("----",err)
//			continue
//		}
//		plg, err := pdll.Lookup("Demo")
//		if err != nil {
//			panic(err)
//		}
//		plg.(func())()
//	}
//}

package gid
//
//import (
//	"fmt"
//	"sync"
//	"testing"
//)
//
////func TestGoID(t *testing.T) {
//func TestGoID(t *testing.T) {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	id0, id1 := Get_gid()
//	go func() {
//		idInternal0, internel1 := Get_gid()
//		fmt.Printf("outer go id : 0->%d 1->%d", id0, id1)
//		fmt.Printf("inner go id: 0->%d 1->%d", idInternal0, internel1)
//		wg.Done()
//	}()
//	wg.Wait()
//}

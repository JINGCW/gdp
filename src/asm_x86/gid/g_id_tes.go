package gid

import (
	"fmt"
	"sync"
)

//func TestGoID(t *testing.T) {
func GoID() {
	var wg sync.WaitGroup
	wg.Add(1)
	//id := Get_gid()
	go func() {
		idInternal := Get_gid()
		fmt.Println(idInternal)
		//assert.NotEqual(t, id, idInternal)
		wg.Done()
	}()
	wg.Wait()
}

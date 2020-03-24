package desync

import (
	"fmt"
	"sync"
	"time"
)

type mulock struct {
	mu sync.Mutex
}

func MuDemo() {
	//m := mulock{}
	//var m sync.Mutex
	//old := reflect.ValueOf(m).FieldByName("state")
	ch := make(chan int)
	//defer m.Unlock()

	//fmt.Println(old)
	go func() {
		ch <- 1
		time.Sleep(3 * time.Second)
		//m.Lock()
		print("go routine1\n")
	}()
	go func() {
		time.Sleep(4 * time.Second)
		//new := old
		//fmt.Println(new)
		ch <- 2
		print("go routine2\n")
	}()

	fmt.Printf("return from channel %d\n",<-ch)
	fmt.Printf("return from channel %d\n",<-ch)
	fmt.Printf("return from final channel %d\n",<-ch)
	sync.Cond.Wait()
}

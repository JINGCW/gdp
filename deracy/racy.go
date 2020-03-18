package deracy

import (
	"fmt"
	"math/rand"
	"time"
)

func TimerRacy() {
	start := time.Now()
	var t *time.Timer

	t = time.AfterFunc(randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			t.Reset(randomDuration())
		})
	time.Sleep(5 * time.Second)
}

func TimerRacySafe() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer

	t = time.AfterFunc(randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			reset <- true
		})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	dura := time.Duration(rand.Int63n(1e9))
	println("dura: ", dura)
	return dura
}

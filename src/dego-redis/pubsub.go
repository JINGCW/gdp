package dego

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func SubPub(addr, passwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	done := make(chan struct{})
	client.Publish("mychannel", "hello!\n")
	go func() {
		pubsub := client.Subscribe("mychannel")
		msg, _ := pubsub.Receive()
		fmt.Println("receive from channel", msg)
		done <- struct{}{}
	}()
	<-done
}

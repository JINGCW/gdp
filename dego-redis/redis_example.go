package dego_redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func ExampleClient(addrp,passwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,
	})
	err := client.Set("qiming", "xie", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("qiming").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("qiming ", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 dose not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2 ", val2)
	}
}

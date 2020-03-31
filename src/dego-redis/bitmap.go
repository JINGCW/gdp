package dego

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func BitMap(addr, passwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})
	fmt.Println(client.GetBit("name", 3).Val())
}

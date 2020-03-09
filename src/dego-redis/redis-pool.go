package dego

import (
	"github.com/go-redis/redis/v7"
	"time"
)

func PoolEX(addr, passwd string) error {

	pool := redis.NewClient(&redis.Options{
		Addr:               addr,
		Password:           passwd,
		DB:                 0,
		MinIdleConns:3,
	})

	for {
		println(pool.RandomKey().Val())
		println("running")
		time.Sleep(3*time.Second)
	}
}

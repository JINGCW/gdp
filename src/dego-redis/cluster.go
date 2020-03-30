package dego

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func Cclient(addr []string, passwd string) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addr,
		Password: passwd,
	})

	fmt.Println(client.Get("name").Val())
}

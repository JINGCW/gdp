package dego

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func Cclient(addr []string, passwd string) error {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addr,
		Password: passwd,
	})
	fmt.Println(client.ClusterInfo())
	println("----")
	//if err := client.Get("name").Err(); err != nil {
	//	return err
	//}
	//fmt.Println(client.Get("name").Val())
	return nil
}

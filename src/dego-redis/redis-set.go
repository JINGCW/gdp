package dego

import "github.com/go-redis/redis/v7"

func SetEx(addr, passwd string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})
	err := client.SAdd("set_ex", "mongdb", "mysql").Err()
	if err != nil {
		return err
	}
	return nil
}

package dego

import "github.com/go-redis/redis/v7"


func ListEX(addrp, ps string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: ps,
		DB:       0,
	})
	err := client.LPush("list_ex", "haha").Err()
	if err != nil {
		return err
	}
	return nil
}

package dego

import "github.com/go-redis/redis/v7"

func HashEX(addrp, passwd string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,})
	client.HMSet("hash_ex", "name", "redis tutorial", "description", "redis basic commands for caching", "likes", 20, "visitors", 2)
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}

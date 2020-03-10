package dego

import "github.com/go-redis/redis/v7"

func ZsetEX(addr, passwd string) error {
	pool := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     passwd,
		DB:           0,
		MinIdleConns: 2,
	})

	out, err := pool.ZAdd("zset_ex", &redis.Z{1, "redis"}, &redis.Z{2, "mongodb"}).Result()
	println(out)
	if err != nil {
		return err
	}
	return nil
}

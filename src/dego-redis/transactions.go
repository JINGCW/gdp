package dego

import "github.com/go-redis/redis/v7"

func Transactions(addr, passwd string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})
	defer client.Close()
	pipe := client.TxPipeline()
	pipe.Set("pipe", 0, 0)
	pipe.Incr("pipe")
	pipe.Incr("pipe")
	pipe.Exec()
	return nil
}

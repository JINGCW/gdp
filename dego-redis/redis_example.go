package dego_redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
)

func ExampleClient(addrp, passwd string) {
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

func ExampleStructSet(addrp, passwd string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,})
	s := struct {
		UserName string
		email    string
		id       int
	}{"otto", "otto@repoment.com", 123456}
	//serialize object to JSON
	js, err := json.Marshal(&s)
	fmt.Println(js)
	if err != nil {
		return err
	}
	e := client.Set("user:"+s.UserName, js, 0).Err()
	if e != nil {
		return e
	}
	return nil
}

func ExampleStructGET(addrp, passwd, key string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,})
	s := struct {
		UserName string
		email    string
		id       int
	}{}
	out,err:=client.Get(key).Bytes()
	if err == redis.Nil {
		fmt.Println("key2 dose not exist")
	} else if err != nil {
		panic(err)
	}
	e := json.Unmarshal(out, &s)
	if e != nil {
		panic(e)
	}

	fmt.Println("bytes: ", out)
	fmt.Printf("unmarshal :%T\n", s)
	fmt.Println(s.UserName)
	fmt.Println(s.email)
	fmt.Println(s.id)
	return nil
}

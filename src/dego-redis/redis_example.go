package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
)

type Usr struct {
	UserName string
	Email    string
	Id       int
}

func ExampleClient(addrp, passwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,
	})
	err := client.Set("hello", "world", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello ", val)

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
	s := Usr{"apple", "otto@repoment.com", 123456}
	//serialize object to JSON
	js, err := json.Marshal(s)
	fmt.Println(js)
	if err != nil {
		return err
	}
	//e := client.Set("user:"+s.UserName, js, 0).Err()
	e := client.Set("user:"+s.UserName,js,0).Err()
	if e != nil {
		return e
	}
	//err= json.Unmarshal(js,&s)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(js)
	return nil
}

func ExampleStructGET(addrp, passwd, key string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addrp,
		Password: passwd,
		DB:       0,})
	s := Usr{}
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

	fmt.Printf("unmarshal :%v\n", s)
	fmt.Println(s.UserName)
	fmt.Println(s.Email)
	fmt.Println(s.Id)
	return nil
}

func ExampleCornerCases(addrp, passwd, key string)error{
	//client := redis.NewClient(&redis.Options{
	//	Addr:     addrp,
	//	Password: passwd,
	//	DB:       0,})

	//custom command
	//res,err:=client.Do("set","key","value").Result()

	return nil;
}

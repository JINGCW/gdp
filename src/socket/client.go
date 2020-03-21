package main

import (
	"fmt"
	"net"
)

func main() {
	var msg = []byte("abcd")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	_, err = conn.Write([]byte(msg))

	fmt.Println("send message: " + string(msg))
	_, err = conn.Read(msg[:])
	fmt.Println("recv message: " + string(msg))
}

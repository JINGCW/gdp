package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go conn_handler(conn)
	}
}

func conn_handler(conn net.Conn) {
	defer conn.Close()
	var body [4]byte

	addr := conn.RemoteAddr()
	for {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("recv %s: %s\n", addr, string(body[:]))
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
	}
	fmt.Printf("disconnected.\n")
}

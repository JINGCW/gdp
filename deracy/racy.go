package deracy

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

func TimerRacy() {
	start := time.Now()
	var t *time.Timer

	t = time.AfterFunc(randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			t.Reset(randomDuration())
		})
	time.Sleep(5 * time.Second)
}

func TimerRacySafe() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer

	t = time.AfterFunc(randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			reset <- true
		})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	dura := time.Duration(rand.Int63n(1e9))
	println("dura: ", dura)
	return dura
}

func SockServer() {
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	if err != nil {
		panic(err)
	}
	defer syscall.Closesocket(sockfd)

	if err = syscall.Bind(sockfd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}); err != nil {
		panic(err)
	}
	println("server running")
	if err = syscall.Listen(sockfd, 3); err != nil {
		panic(err)
	}
	//println("server running")

	var d syscall.WSAData
	syscall.WSAStartup(uint32(0x202), &d)
	var buf []byte
	bufdata := syscall.WSABuf{
		Len: uint32(4),
		Buf: &buf[0],
	}
	o := syscall.Overlapped{}
	flags := uint32(0)
	qty := uint32(0)
	for {
		syscall.Accept(sockfd)
		//bufdata := syscall.WSABuf{
		//	Len: uint32(4),
		//	Buf: &buf[0],
		//}
		//o := syscall.Overlapped{}
		//flags := uint32(0)
		//qty := uint32(0)
		syscall.WSARecv(sockfd, &bufdata, 1, &qty, &flags, &o, nil)
		//syscall.Recvfrom(sockfd, buf, 1)
		//fmt.Println(bufdata)
		if len(buf) != 0 {
			println(buf)
		}
	}
}

func SockClient() {
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	if err != nil {
		panic(err)
	}
	defer syscall.Closesocket(sockfd)

	syscall.Connect(sockfd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1}},
	)
	var (
		data       syscall.WSABuf
		sendstr    string = "hello"
		SendButes  uint32
		overlapped syscall.Overlapped
	)
	data.Len = uint32(len(sendstr))
	data.Buf, err = syscall.BytePtrFromString(sendstr)
	if err != nil {
		panic(err)
	}
	//如果使用syscall.Sendto或syscall.Write会发送失败，原因未知
	err = syscall.WSASend(sockfd, &data, 1, &SendButes, 0, &overlapped, nil)
	if err != nil {
		println("Send error")
	} else {
		println("Send success")
	}
}

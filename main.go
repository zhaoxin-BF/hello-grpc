package main

import (
	"fmt"
	person1 "github.com/zhaoxin-BF/hello-grpc/golang/person/v1"
	"github.com/zhaoxin-BF/hello-grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// 一、
func main() {
	rpcServer := grpc.NewServer()

	person1.RegisterPersonServiceServer(rpcServer, service.PersonService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}

	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动监听出错", err)
	}
	fmt.Println("启动grpc服务成功！")

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-signals:
		fmt.Println("terminating via signal")
	}

	return
}

// 二、序列化与反序列化
//func () {
//	passWord := "123456"
//
//	person := &person2.Person{
//		Age:      12,
//		Name:     "boreas",
//		Password: &passWord,
//		Hobbies:  []string{"游泳", "健身"},
//	}
//
//	// 序列化
//	marshal, err := proto.Marshal(person)
//	if err != nil {
//		panic(err)
//	}
//
//	// 反序列化
//	newPerson := &person2.Person{}
//
//	err = proto.Unmarshal(marshal, newPerson)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(newPerson)
//}

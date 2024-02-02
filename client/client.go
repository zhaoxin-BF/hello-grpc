package main

import (
	"context"
	"fmt"
	"github.com/zhaoxin-BF/hello-grpc/golang/person/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("服务端出错！")
	}

	defer conn.Close()

	resdDirs()

	personClient := person.NewPersonServiceClient(conn)
	request := &person.GetPersonRequest{PersonName: "boreas"}

	personResp, err := personClient.GetPerson(context.Background(), request)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("查询失败", personResp)
	} else {
		fmt.Println("查询成功", personResp)
	}

}

func resdDirs() {
	dir, err := os.Open("../proto")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(dir)

	defer dir.Close()
	fileInfos, _ := dir.Readdir(-1)

	fmt.Println(fileInfos)
	for _, v := range fileInfos {
		if v.IsDir() {
			fmt.Println(v.Name())
			continue
		}
		fmt.Println(v)
	}
}

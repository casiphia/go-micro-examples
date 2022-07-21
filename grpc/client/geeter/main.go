package main

import (
	"context"
	"fmt"
	"geeter/proto/geete"
	"geeter/proto/geeter"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"google.golang.org/grpc"
)

func main() {
	microCall()

	grpcCall()
}

func microCall() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.client.geeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	cli := geeter.NewGeeterService("go.micro.service.geeter", client.DefaultClient)
	resp, err := cli.Call(context.TODO(), &geeter.Request{
		Name: "micro call.",
	})
	fmt.Println(resp, err)
}

func grpcCall() {

	// 创建客户端连接
	conn, err := grpc.Dial("127.0.0.1:53294", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc Dial err: ", err)
		return
	}
	defer conn.Close()
	// 创建远程服务的客户端
	cli := geete.NewGeeterClient(conn)
	resp, err := cli.Call(context.TODO(), &geete.Request{
		Name: "micro call.",
	})
	fmt.Println(resp, err)

}

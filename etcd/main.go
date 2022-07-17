package main

/*
 * @Descripttion: 828 project
 * @Version: 1.0.0
 * @Author: wangxiaodiao
 * @Email: 413586280@qq.com
 * @Date: 2021-11-01 18:52:17
 * @LastEditors: wangxiaodiao
 * @LastEditTime: 2022-07-17 22:18:23
 * @FilePath: /vcg-global/app/auctionmgr/main.go
 */

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)), //etcd注册
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

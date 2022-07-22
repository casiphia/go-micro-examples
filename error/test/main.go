package main

import (
	"test/handler"
	"test/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	test "test/proto/test"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.test"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.test", service.Server(), new(subscriber.Test))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

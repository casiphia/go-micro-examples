package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"geeter/handler"
	"geeter/subscriber"

	geeter "geeter/proto/geeter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.geeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	geeter.RegisterGeeterHandler(service.Server(), new(handler.Geeter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.geeter", service.Server(), new(subscriber.Geeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

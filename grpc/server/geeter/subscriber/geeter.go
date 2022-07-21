package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	geeter "geeter/proto/geeter"
)

type Geeter struct{}

func (e *Geeter) Handle(ctx context.Context, msg *geeter.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *geeter.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}

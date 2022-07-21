package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	geeter "geeter/proto/geeter"
)

type Geeter struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Geeter) Call(ctx context.Context, req *geeter.Request, rsp *geeter.Response) error {
	log.Info("Received Geeter.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Geeter) Stream(ctx context.Context, req *geeter.StreamingRequest, stream geeter.Geeter_StreamStream) error {
	log.Infof("Received Geeter.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&geeter.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Geeter) PingPong(ctx context.Context, stream geeter.Geeter_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&geeter.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

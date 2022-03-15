package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	mymicro "mymicro/proto"
)

type Mymicro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Mymicro) Call(ctx context.Context, req *mymicro.Request, rsp *mymicro.Response) error {
	log.Info("Received Mymicro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Mymicro) Stream(ctx context.Context, req *mymicro.StreamingRequest, stream mymicro.Mymicro_StreamStream) error {
	log.Infof("Received Mymicro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&mymicro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Mymicro) PingPong(ctx context.Context, stream mymicro.Mymicro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&mymicro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

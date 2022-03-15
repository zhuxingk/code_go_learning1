package main

import (
	"mymicro/handler"
	pb "mymicro/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("mymicro"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMymicroHandler(srv.Server(), new(handler.Mymicro))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

package main

import (
	"net"

	"github.com/itouri/petitVDC/api"
)

func main() {
	grpcListener, err := net.Listen("tcp", "127.0.0.1:5050")
	if err != nil {
		panic(err)
	}
	grpcServer := api.NewAPIServer()
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			panic(err)
		}
	}()
	defer grpcServer.GracefulStop()
}

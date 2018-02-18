package api

import (
	"google.golang.org/grpc"
)

type APIServer struct {
	server *grpc.Server
}

type InstanceAPIServer struct {
	api *APIServer
}

func NewAPIServer() *APIServer {
	s := &APIServer{
		server: grpc.NewServer(),
	}
	s.server.RegisterService(&_Instance_serviceDesc, &InstanceAPIServer{api: s})
	return s
}

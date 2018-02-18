package api

import (
	"net"

	"google.golang.org/grpc"
)

type APIServer struct {
	server *grpc.Server
}

func NewAPIServer() *APIServer {
	s := &APIServer{
		server: grpc.NewServer(),
	}
	s.server.RegisterService(&_Instance_serviceDesc, &InstanceAPI{api: s})
	return s
}

func (s *APIServer) Serve(listen net.Listener) error {
	return s.server.Serve(listen)
}

func (s *APIServer) GracefulStop() {
	s.server.GracefulStop()
}

package util

import "google.golang.org/grpc"

func RemoteCall(c func(*grpc.ClientConn) error) error {
	serverAddr := "127.0.0.1:5050" // TODO changing to refering config fite
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	return c(conn)
}

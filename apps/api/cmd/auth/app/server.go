package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewAuthServer() func() int {
	
	return func() int {
		return 1
	}
}



func createServerChai() error {
	// 启动服务
	srv := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = srv.Serve(lis)

	return err
}

package main

import (
	pb "gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8089"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

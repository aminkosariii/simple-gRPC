package main

import (
	"context"
	pb "gRPC/proto"
)

func (s *helloServer) sayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hiiii",
	}, nil
}

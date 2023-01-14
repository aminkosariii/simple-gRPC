package main

import (
	pb "gRPC/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with these names: %s", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "this is response message for this name" + name,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

package main

import (
	pb "gRPC/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}

		log.Fatalf("Got request with this name : %v", request.Name)
		messages = append(messages, "Hellooo"+request.Name)
	}
}

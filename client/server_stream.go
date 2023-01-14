package main

import (
	"context"
	pb "gRPC/proto"
	"io"
	"log"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("stream started...")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not stream with this error %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished")

}

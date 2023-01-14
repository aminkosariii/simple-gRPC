package main

import (
	"context"
	"fmt"
	pb "gRPC/proto"
	"log"
	"time"
)

func CallSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {

	fmt.Printf("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send stream from client side: %v", err)
	}
	for _, name := range names.Names {
		request := &pb.HelloRequest{
			Name: name,
		}
		err := stream.Send(request)
		if err != nil {
			log.Fatalf("can not send names with this error: %v", err)
		}
		log.Printf("send request with this name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiveing %v", err)
	}
	log.Printf("%v", res.Messages)
}

package main

import (
	"context"
	"log"
	"time"
	pb "grpc-with-go/proto"	
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not able to send names: %v", err)
	}

	//sending names stream to the server
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent request's with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished!")
	if err != nil {
		log.Fatalf("Error while receiving ack from server: %v", err)
	}
	log.Printf("%v\n", res.Messages)
}

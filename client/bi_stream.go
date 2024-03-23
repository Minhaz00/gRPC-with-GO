package main

import (
	"context"
	"io"
	"log"
	"time"
	pb "grpc-with-go/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming has been started")
	stream, err := client.SayHelloBidiStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	//making a channel
	waitc := make(chan struct{})

	//go routine
	go func()  {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest {
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished.")
}

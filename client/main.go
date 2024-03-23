package main

import (
	pb "grpc-with-go/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	//connection close
	defer conn.Close()

	//creating a new client
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Minhaz", "Jisun", "Alvi"},
	}

	// callSaySayHello(client) //unart api
	// callSayHelloServerStream(client, names) // server streaming
	// callSayHelloClientStream(client, names) // client streaming
	callSayHelloBidirectionalStream(client, names) // bidirectional streaming

}

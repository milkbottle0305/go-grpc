package main

import (
	"context"
	"fmt"
	pb "grpc/client-streaming/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func make_message(message string) pb.Message {
	return pb.Message{
		Message: message,
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewBidrectionalClient(conn)
	stream, err := c.GetServerResponse(context.Background())
	if err != nil {
		log.Fatalf("Opening stream", err)
	}
	messages := []pb.Message{
		make_message("message #1"),
		make_message("message #2"),
		make_message("message #3"),
		make_message("message #4"),
		make_message("message #5"),
	}
	for _, message := range messages {
		if err := stream.Send(&message); err != nil {
			log.Fatalln("Send", err)
		}
		fmt.Println("[client to server]", message.GetMessage())
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}
		fmt.Println("[server to client]", res.GetMessage())
	}
}

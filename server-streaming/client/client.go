package main

import (
	"context"
	"fmt"
	pb "grpc/server-streaming/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func recv_message(c pb.ServerStreamingClient) {
	req := &pb.Number{
		Value: 5,
	}
	stream, err := c.GetServerResponse(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while Calling GetServerReseponse", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while reciving stream", err)
		}
		fmt.Println("[server to client]", res.GetMessage())
	}
}

func main() {
	// 기본적으로는 ssl init을 지원하지만, ssl 작업 없이 시작하기 위해 insecure 옵션 사용
	// 1. create the connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	// conn 커넥션이 프로그램 종료 시 같이 끊기도록 설정함. defer는 프로그램 말미에 실행되도록 하는 명령어
	// 3. close connection. (end of program)
	defer conn.Close()

	// 2. create a client
	c := pb.NewServerStreamingClient(conn)
	recv_message(c)
}

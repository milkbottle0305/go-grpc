package main

import (
	"context"
	"fmt"
	pb "grpc/client-streaming/proto"
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
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //서버와 insecure방식으로 연결
	if err != nil {                                                                                     // 연결과정에서 에러 핸들링
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close() // 프로그램 종료전 서버 닫기

	c := pb.NewClientStreamingClient(conn)                   // 클라이언트 생성
	stream, err := c.GetServerResponse(context.Background()) // 서버에게 요청

	if err != nil { // 에러 발생 시 핸들링
		log.Fatalf("Opening stream", err)
	}
	messages := []pb.Message{ // 메시지 5개를 생성
		make_message("message #1"),
		make_message("message #2"),
		make_message("message #3"),
		make_message("message #4"),
		make_message("message #5"),
	}
	for _, message := range messages { // 반복문을 통해 하나씩 요청
		if err := stream.Send(&message); err != nil { // stream을 요청하다가 에러 발생시
			log.Fatalln("Send", err)
		}
		fmt.Println("[client to server]", message.GetMessage()) // 요청한 stream에 대한 message를 get(송신)
	}
	if err != nil {
		log.Fatalln("Error while Calling GetServerResponse RPC : ", err)
	}
	res, err := stream.CloseAndRecv()
	fmt.Println("[server to client]", res.GetValue())
}

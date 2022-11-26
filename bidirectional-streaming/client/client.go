package main

import (
	"context"
	"fmt"
	pb "grpc/bidirectional-streaming/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func make_message(message string) pb.Message { // 메시지 만들기
	return pb.Message{
		Message: message,
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // 서버와의 연결
	if err != nil {                                                                                     // 연결 과정 중 에러 발생시
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close() // 프로그램 종료 전 서버 종료

	c := pb.NewBidrectionalClient(conn)                      // 클라리언트 생성
	stream, err := c.GetServerResponse(context.Background()) // 서버에게 요청
	if err != nil {
		log.Fatalf("Opening stream", err)
	}
	messages := []pb.Message{ // 메시지 5개를 생성
		make_message("message #1"),
		make_message("message #2"),
		make_message("message #3"),
		make_message("message #4"),
		make_message("message #5"),
	}
	for _, message := range messages { // 반복문을 통해 하나씩 서버에게 요청
		if err := stream.Send(&message); err != nil { // 서버에게 요청, 도중 에러발생시 핸들링
			log.Fatalln("Send", err)
		}
		fmt.Println("[client to server]", message.GetMessage())
	}
	if err := stream.CloseSend(); err != nil { // 요청을 다 보냈으면 stream을 닫음
		log.Fatalln("CloseSend", err)
	}
	for { // 요청에 대한 응답 또한 stream방식이므로 반복
		res, err := stream.Recv() // stream인 응답을 res에 저장
		if err == io.EOF {        // for문 종료 조건, 끝까지가면 반복문 종료
			break
		}
		if err != nil { // 응답을 받다가 에러 발생시 핸들링
			log.Fatalln("Recv", err)
		}
		fmt.Println("[server to client]", res.GetMessage()) // 요청에 대한 결과를 출력
	}
}

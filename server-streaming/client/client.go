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
	req := &pb.Number{ // 몇 번의 송신을 받는지 요청을 나타내는 req 생성
		Value: 5,
	}
	stream, err := c.GetServerResponse(context.Background(), req) // 서버에게 요청
	if err != nil {
		log.Fatalln("Error while Calling GetServerReseponse", err)
	}
	for { // server streaming이므로 서버에게 받을 stream만큼 반복
		res, err := stream.Recv() // 서버에게서 stream을 수신
		if err == io.EOF {        // 끝까지 받으면 종료
			break
		}
		if err != nil { // 받는 도중 에러 발생시 핸들링
			log.Fatalln("Error while reciving stream", err)
		}
		fmt.Println("[server to client]", res.GetMessage()) // 응답인 메시지를 받으며 콘솔 출력
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

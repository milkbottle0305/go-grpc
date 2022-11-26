package main

import (
	"fmt"
	pb "grpc/client-streaming/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedClientStreamingServer
}

func (*server) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	var count int32 = 0 // 몇번의 요청을 받았는지 저장하는 변수 정의
	for {               // 요청이 stream 방식이므로 반복
		_, err := stream.Recv() // 요청 stream을 수신
		if err == io.EOF {      // stream을 끝까지 받았다면
			return stream.SendAndClose(&pb.Number{ // 요청 횟수를 나타내는 count를 반환 후 수신 stream을 닫음
				Value: count,
			})
		}
		if err != nil { // 수신 하다가 에러 발생시 핸들링
			log.Fatalln("Recv", err)
		}
		count += 1 // 요청 횟수를 1 증가
	}
}

func main() {
	fmt.Println("Starting server. Listening on port 50051.")

	// 50051 TCP방식으로 서버 ON
	lis, err := net.Listen("tcp", "[0.0.0.0]:50051")
	if err != nil {
		log.Fatalf("Failed to listen", err)
	}

	// 2. grpc server open
	s := grpc.NewServer()
	pb.RegisterClientStreamingServer(s, &server{})

	// port와 grpc server 연결 후 serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

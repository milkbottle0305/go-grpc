package main

import (
	"fmt"
	pb "grpc/bidirectional-streaming/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBidrectionalServer
}

func (*server) GetServerResponse(stream pb.Bidrectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC bidirectional streaming.")
	for { //클라이언트의 요청도 stream이고, 응답도 stream이므로 반복
		req, err := stream.Recv() // req에 stream을 수신
		if err == io.EOF {        // stream의 끝을 나타내는 EOF에 간다면 종료
			return nil
		}
		if err != nil { // 수신하는 과정에서 에러가 발생할 때
			log.Fatalln("Recv", err)
		}
		res := stream.Send(&pb.Message{ // res로 stream을 송신
			Message: req.GetMessage(),
		})
		if res != nil { // 송신하는 과정에서 에러가 발생할 때
			log.Fatalf("Error when response was sent to the client: %v", res)
		}
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
	pb.RegisterBidrectionalServer(s, &server{})

	// port와 grpc server 연결 후 serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

package main

import (
	"context"
	"fmt"
	"grpc/unary/hello_grpc"
	pb "grpc/unary/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) MyFunction(ctx context.Context, req *pb.MyNumber) (*pb.MyNumber, error) {
	result := hello_grpc.My_func(req.GetValue())
	// 리턴할 Response 객체 생성. (포인터 response이므로 &를 객체 앞에 붙여준다.)
	res := &pb.MyNumber{
		Value: result,
	}
	return res, nil
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
	pb.RegisterMyServiceServer(s, &server{})

	// port와 grpc server 연결 후 serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

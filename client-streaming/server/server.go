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
	pb.UnimplementedBidrectionalServer
}

func (*server) GetServerResponse(stream pb.Bidrectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	var count int32 = 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Number{
				Value: count,
			})
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}
		count += 1
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

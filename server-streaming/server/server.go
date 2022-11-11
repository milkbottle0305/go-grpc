package main

import (
	"fmt"
	pb "grpc/server-streaming/proto"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

func make_message(message string) pb.Message {
	return pb.Message{
		Message: message,
	}
}

type server struct{}

func (*server) GetServerResponse(req *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	var message []pb.Message
	for i := 0; i < int(req.Value); i++ {
		message = append(message, make_message("message #"+strconv.Itoa(i+1)))
	}
	fmt.Println("Server processing gRPC server-streaming" + strconv.Itoa(int(req.Value)) + ".")
	for i := 0; i < int(req.Value); i++ {
		stream.Send(&message[i])
	}
	return nil
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
	pb.RegisterServerStreamingServer(s, &server{})

	// port와 grpc server 연결 후 serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

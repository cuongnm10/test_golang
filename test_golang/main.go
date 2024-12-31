package main

import (
	"context"
	"log"
	"net"

	pb "github.com/cuongnm10/test_golang/personpb"

	"google.golang.org/grpc"
)

// Server struct để implement dịch vụ gRPC
type Server struct {
	pb.UnimplementedPersonServiceServer
}

// Hàm GetPerson để xử lý request
func (s *Server) GetPerson(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	log.Printf("Received request for ID: %d", req.Id)

	// Trả về dữ liệu giả lập
	return &pb.PersonResponse{
		Id:    req.Id,
		Name:  "Nguyen Van A",
		Age:   30,
		Email: "nguyenvana@example.com",
	}, nil
}

func main() {
	// Lắng nghe trên cổng 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Tạo server gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, &Server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

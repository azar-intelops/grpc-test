package main

import (
	"context"
	"log"
	"net"

	pb "github.com/azar-intelops/go-interceptors/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Define a gRPC interceptor
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Received request: %v", req)
	resp, err := handler(ctx, req)
	return resp, err
}

type Server struct {
	pb.UnimplementedMyServiceServer
}

func (s *Server) DemoMethod(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	return &pb.DemoResponse{
		Message: "Hello " + req.Message,
	}, nil
}

func main() {
	// Create a new gRPC server with the logging interceptor
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	// Register your gRPC service with the server
	myService := &Server{}
	pb.RegisterMyServiceServer(s, myService)
	reflection.Register(s)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting server in port :%d\n", 50051)

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

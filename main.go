package main

import (
	"context"
	"log"
	"net"

	"github.com/azar-intelops/go-interceptors/configs"
	"github.com/azar-intelops/go-interceptors/controllers"
	pb "github.com/azar-intelops/go-interceptors/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Define a gRPC interceptor
func loggingInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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

	db := configs.DB
	// log.Println(db)

	// Register your gRPC service with the server
	myService := &Server{}
	authService := controllers.NewAuthServer(db)
	authserver := controllers.NewAuthServiceServer(db)
	pb.RegisterMyServiceServer(s, myService)
	pb.RegisterUserServiceServer(s, authService)
	pb.RegisterAuthServiceServer(s, authserver)
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

package controllers

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/azar-intelops/go-interceptors/configs"
	"github.com/azar-intelops/go-interceptors/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufsize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufsize)
	s := grpc.NewServer()
	db := configs.DB
	authServer := NewAuthServer(db)
	pb.RegisterUserServiceServer(s, authServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	resp, err := client.CreateUser(ctx, &pb.UserRequest{
		Name:   "Ram",
		Mobile: 8958589656,
		Email:  "test@compage.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetStatus() != "user successfully added!" {
		t.Fatal("User not created!")
	}
}

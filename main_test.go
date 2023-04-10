package main

import (
	"context"
	"log"
	"net"
	"testing"

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
	pb.RegisterMyServiceServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestDemoMethod(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewMyServiceClient(conn)
	// Case 1
	resp, err := client.DemoMethod(ctx, &pb.DemoRequest{Message: "world"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetMessage() != "Hello world" {
		t.Fatal("Demo Response must be 'Hello world'")
	}
	// Case 2
	resp1, err := client.DemoMethod(ctx, &pb.DemoRequest{Message: "Compage"})
	if err != nil {
		t.Fatal(err)
	}
	if resp1.GetMessage() != "Hello Compage" {
		t.Fatal("Demo Response must be 'Hello Compage'")
	}

}

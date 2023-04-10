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

func demoHelper(text string, ctx context.Context, client pb.MyServiceClient, t *testing.T) {
	resp, err := client.DemoMethod(ctx, &pb.DemoRequest{Message: text})
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetMessage() != "Hello "+text {
		t.Fatal("Demo Response must be 'Hello " + text + "'")
	}
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
	demoHelper("world", ctx, client, t)
	// Case 2
	demoHelper("Compage", ctx, client, t)

}

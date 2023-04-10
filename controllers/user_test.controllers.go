package controllers

// import (
// 	"context"
// 	"log"

// 	"github.com/azar-intelops/go-interceptors/configs"
// 	"github.com/azar-intelops/go-interceptors/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/test/bufconn"
// )

// func server(ctx context.Context) (pb.MyServiceClient, func()) {
// 	buffer := 101024 * 1024
// 	lis := bufconn.Listen(buffer)
// 	baseServer := grpc.NewServer()
// 	db := configs.DB
// 	pb.RegisterUserServiceServer(baseServer, NewAuthServer(db))
// 	go func() {
// 		if err := baseServer.Serve(lis); err != nil {
// 			log.Printf("error serving server: %v", err)
// 		}
// 	}()
// }

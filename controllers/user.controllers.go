package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/azar-intelops/go-interceptors/configs"
	"github.com/azar-intelops/go-interceptors/models"
	"github.com/azar-intelops/go-interceptors/pb"
	"github.com/azar-intelops/go-interceptors/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	pb.UnimplementedUserServiceServer
	client *mongo.Client
}

func NewAuthServer(client *mongo.Client) *AuthServer {
	return &AuthServer{
		client: client,
	}
}
func (s *AuthServer) GetUser(ctx context.Context, req *pb.EmptyMessage) (*pb.UserResponses, error) {
	userCollection := configs.GetCollection(s.client, "users")
	var results []models.UserReq
	cur, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Found multiple documents: %+v\n", results)
	var response []*pb.UserRequest
	for _, value := range results {
		response = append(response, &pb.UserRequest{
			Id:        value.Id.Hex(),
			Name:      value.Name,
			Mobile:    value.Mobile,
			Email:     value.Email,
			CreatedAt: timestamppb.New(value.CreatedAt),
			UpdatedAt: timestamppb.New(value.UpdatedAt),
		})

	}
	fmt.Println(response)
	//Close the cursor once finished
	defer cur.Close(context.TODO())

	return &pb.UserResponses{
		Req: response,
	}, nil
}
func (s *AuthServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if utils.IsEmpty(req.GetName()) || utils.IsEmpty(req.GetEmail()) || req.GetMobile() == 0 {
		return nil, errors.New("values cant be empty")
	}

	userCollection := configs.GetCollection(s.client, "users")
	fmt.Println(userCollection)
	user := models.UserReq{
		Id:        primitive.NewObjectID(),
		Name:      req.GetName(),
		Mobile:    req.GetMobile(),
		Email:     req.GetEmail(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID)
	return &pb.UserResponse{
		Status: "user successfully added!",
	}, nil
}

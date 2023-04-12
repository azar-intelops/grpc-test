package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/azar-intelops/go-interceptors/configs"
	"github.com/azar-intelops/go-interceptors/models"
	"github.com/azar-intelops/go-interceptors/pb"
	"github.com/azar-intelops/go-interceptors/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	client *mongo.Client
}

func NewAuthServiceServer(client *mongo.Client) *AuthServiceServer {
	return &AuthServiceServer{
		client: client,
	}
}

func validateEmpty(values ...interface{}) bool {
	for _, value := range values {
		if utils.IsEmptyAnyType(value) {
			return true
		}
	}
	return false
}

func (s *AuthServiceServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	if validateEmpty(req.GetName(), req.GetMobile(), req.GetEmail(), req.GetPassword()) {
		return nil, errors.New("values can't be empty, please try again")
	}
	authCollection := configs.GetCollection(s.client, "users")
	user := models.SignupRequest{
		Id:        primitive.NewObjectID(),
		Name:      req.GetName(),
		Mobile:    req.GetMobile(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := authCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("something went wrong while inserting response, try again")
	}
	return &pb.SignupResponse{
		Id: res.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

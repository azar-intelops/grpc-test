package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupRequest struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Mobile    int64              `bson:"mobile"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserReq struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Mobile    int64              `bson:"mobile"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

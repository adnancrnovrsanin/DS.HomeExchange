package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Home struct {
	Id                       primitive.ObjectID `json:"id,omitempty" bson:"home_id,omitempty"`
	OwnerId                  string             `json:"ownerId" bson:"ownerId"`
	Address                  string             `json:"address" bson:"address"`
	Description              string             `json:"description" bson:"description"`
	NumberOfRequestsRejected int                `json:"numberOfRequestsRejected" bson:"numberOfRequestsRejected"`
}

type CreateHomeRequest struct {
	OwnerId     string `json:"ownerId" bson:"ownerId"`
	Address     string `json:"address" bson:"address"`
	Description string `json:"description" bson:"description"`
}

type SuccesfulHomeExchangeRequest struct {
	FromUserId     string `json:"fromUserId" bson:"fromUserId"`
	ToUserId       string `json:"toUserId" bson:"toUserId"`
	FromUserHomeId string `json:"fromUserHomeId" bson:"fromUserHomeId"`
	ToUserHomeId   string `json:"toUserHomeId" bson:"toUserHomeId"`
}

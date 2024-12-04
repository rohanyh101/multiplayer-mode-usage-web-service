package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoMode struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type MongoModeTraffic struct {
	ModeID primitive.ObjectID `bson:"mode_id"`
	Users  int32              `bson:"users"`
}

type MongoAreaCode struct {
	ID          primitive.ObjectID `bson:"_id"`
	AreaCode    string             `bson:"area_code"`
	ModeTraffic []MongoModeTraffic `bson:"mode_traffic"`
}

type CacheMode struct {
	Name  string `json:"name"`
	Users int32  `json:"users"`
}

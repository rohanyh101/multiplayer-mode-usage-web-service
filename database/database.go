package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/roohanyh/lila_p1/config"
	"github.com/roohanyh/lila_p1/models"
	"github.com/roohanyh/lila_p1/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	database               = config.Env.MONGO_DB
	modeCollectionName     = config.Env.MONGO_MODE_COLL
	areaCodeCollectionName = config.Env.MONGO_AREACODE_COLL
)

var modeCollection *mongo.Collection = OpenCollection(database, modeCollectionName)
var areaCodeCollection *mongo.Collection = OpenCollection(database, areaCodeCollectionName)

func GetTopMode(ac string) (*proto.Mode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"area_code": ac}
	var mongoAreaCode models.MongoAreaCode
	err := areaCodeCollection.FindOne(ctx, filter).Decode(&mongoAreaCode)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("area code %s not found", ac)
		}
		return nil, err
	}

	if len(mongoAreaCode.ModeTraffic) == 0 {
		return nil, fmt.Errorf("no modes found for area code %s", ac)
	}

	var topTrafficMode models.MongoModeTraffic
	for _, traffic := range mongoAreaCode.ModeTraffic {
		if traffic.Users > topTrafficMode.Users {
			topTrafficMode = traffic
		}
	}

	var mongoMode models.MongoMode
	filter = bson.M{"_id": topTrafficMode.ModeID}
	err = modeCollection.FindOne(ctx, filter).Decode(&mongoMode)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("mode with ID %s not found", topTrafficMode.ModeID.Hex())
		}
		return nil, err
	}

	return &proto.Mode{
		Name:  mongoMode.Name,
		Users: topTrafficMode.Users,
	}, nil
}

func GetModeByName(mn string) (models.MongoMode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": mn}
	var mongoMode models.MongoMode
	err := modeCollection.FindOne(ctx, filter).Decode(&mongoMode)
	if err != nil {
		return models.MongoMode{}, err
	}

	return mongoMode, nil
}

func GetAreaCode(ac string) (models.MongoAreaCode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"area_code": ac}
	var mongoAreaCode models.MongoAreaCode
	err := areaCodeCollection.FindOne(ctx, filter).Decode(&mongoAreaCode)
	if err != nil {
		return models.MongoAreaCode{}, err
	}

	return mongoAreaCode, nil
}

func UpdateSingleMode(ac string, mID primitive.ObjectID, users int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"area_code": ac, "mode_traffic.mode_id": mID}
	update := bson.M{"$set": bson.M{"mode_traffic.$.users": users}}
	_, err := areaCodeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update mode usage for area code %s and mode ID %s: %w", ac, mID.Hex(), err)
	}

	return nil
}

func UpdateModeTraffic(ac models.MongoAreaCode) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"area_code": ac.AreaCode}
	update := bson.M{"$set": bson.M{"mode_traffic": ac.ModeTraffic}}
	_, err := areaCodeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update mode traffic for area code %s: %w", ac.AreaCode, err)
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/roohanyh/lila_p1/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	MONGO_URI := fmt.Sprintf("mongodb://%s:%s@localhost:27017", config.Env.MONGO_USER, config.Env.MONGO_PASS)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	db := client.Database("multiplayer")

	modesColl := db.Collection("modes")
	mode1 := bson.M{"name": "Battle Royale"}
	mode2 := bson.M{"name": "Team Deathmatch"}
	mode3 := bson.M{"name": "Capture the Flag"}

	res1, err := modesColl.InsertOne(context.TODO(), mode1)
	if err != nil {
		log.Fatalf("Failed to insert mode1: %v", err)
	}

	res2, err := modesColl.InsertOne(context.TODO(), mode2)
	if err != nil {
		log.Fatalf("Failed to insert mode2: %v", err)
	}

	res3, err := modesColl.InsertOne(context.TODO(), mode3)
	if err != nil {
		log.Fatalf("Failed to insert mode3: %v", err)
	}

	areaCodesColl := db.Collection("area_codes")
	areaCodes := []interface{}{
		bson.M{
			"area_code": "123",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 500},
				{"mode_id": res2.InsertedID, "users": 300},
				{"mode_id": res3.InsertedID, "users": 200},
			},
		},
		bson.M{
			"area_code": "456",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 600},
				{"mode_id": res2.InsertedID, "users": 400},
				{"mode_id": res3.InsertedID, "users": 300},
			},
		},
		bson.M{
			"area_code": "789",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 700},
				{"mode_id": res2.InsertedID, "users": 250},
				{"mode_id": res3.InsertedID, "users": 100},
			},
		},
		bson.M{
			"area_code": "101",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 450},
				{"mode_id": res2.InsertedID, "users": 350},
				{"mode_id": res3.InsertedID, "users": 200},
			},
		},
		bson.M{
			"area_code": "202",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 800},
				{"mode_id": res2.InsertedID, "users": 500},
				{"mode_id": res3.InsertedID, "users": 300},
			},
		},
		bson.M{
			"area_code": "303",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 400},
				{"mode_id": res2.InsertedID, "users": 450},
				{"mode_id": res3.InsertedID, "users": 150},
			},
		},
		bson.M{
			"area_code": "404",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 550},
				{"mode_id": res2.InsertedID, "users": 200},
				{"mode_id": res3.InsertedID, "users": 250},
			},
		},
		bson.M{
			"area_code": "505",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 300},
				{"mode_id": res2.InsertedID, "users": 600},
				{"mode_id": res3.InsertedID, "users": 400},
			},
		},
		bson.M{
			"area_code": "606",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 700},
				{"mode_id": res2.InsertedID, "users": 500},
				{"mode_id": res3.InsertedID, "users": 600},
			},
		},
		bson.M{
			"area_code": "707",
			"mode_traffic": []bson.M{
				{"mode_id": res1.InsertedID, "users": 200},
				{"mode_id": res2.InsertedID, "users": 300},
				{"mode_id": res3.InsertedID, "users": 100},
			},
		},
	}

	_, err = areaCodesColl.InsertMany(context.TODO(), areaCodes)
	if err != nil {
		log.Fatalf("Failed to insert area codes: %v", err)
	}

	log.Println("data inserted successfully...")
}

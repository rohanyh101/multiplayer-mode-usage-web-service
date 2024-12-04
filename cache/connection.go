package cache

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/roohanyh/lila_p1/config"
)

var dbClient *redis.Client

func Init() {
	dbClient = redis.NewClient(&redis.Options{
		Addr:     config.Env.REDIS_ADDR,
		Password: config.Env.REDIS_PASS,
		DB: func() int {
			dbNumber, err := strconv.Atoi(config.Env.REDIS_DB)
			if err != nil {
				log.Fatalf("Failed to convert REDIS_DB to int: %v", err)
			}

			return dbNumber
		}(),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	log.Println("Connected to redis server successfully")
}

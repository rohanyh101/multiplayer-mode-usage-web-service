package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT                string
	ENV                 string
	REDIS_ADDR          string
	REDIS_PASS          string
	REDIS_DB            string
	MONGO_HOST          string
	MONGO_PORT          string
	MONGO_USER          string
	MONGO_PASS          string
	MONGO_DB            string
	MONGO_MODE_COLL     string
	MONGO_TRAFFIC_COLL  string
	MONGO_AREACODE_COLL string
}

var Env = initConfig()

func initConfig() *Config {

	if os.Getenv("ENV") != "development" {
		log.Printf("using production variables config/config.go")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		log.Printf("using development vars config/config.go")
	}

	return &Config{
		PORT:                getEnv("PORT", "50051"),
		ENV:                 getEnv("ENV", "production"),
		REDIS_ADDR:          getEnv("REDIS_ADDR", "redis:6379"),
		REDIS_PASS:          getEnv("REDIS_PASS", "nopass"),
		REDIS_DB:            getEnv("REDIS_DB", "0"),
		MONGO_USER:          getEnv("MONGO_USER", "root"),
		MONGO_PASS:          getEnv("MONGO_PASSWORD", "nopass"),
		MONGO_HOST:          getEnv("MONGO_HOST", "mongo"),
		MONGO_PORT:          getEnv("MONGO_PORT", "27017"),
		MONGO_DB:            getEnv("DB_NAME", "multiplayer"),
		MONGO_MODE_COLL:     getEnv("MONGO_MODE_COLL", "modes"),
		MONGO_TRAFFIC_COLL:  getEnv("MONGO_TRAFFIC_COLL", "mode_traffic"),
		MONGO_AREACODE_COLL: getEnv("MONGO_AREACODE_COLL", "area_codes"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

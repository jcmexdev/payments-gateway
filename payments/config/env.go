package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type AppConfig struct {
	AppPort       string `env:"API_PORT,default=8080"`
	AppBaseUrl    string `env:"API_URL,required"`
	MongoUrl      string `env:"MONGO_URL, required"`
	MongoDatabase string `env:"MONGO_DATABASE, required"`
}

var Conf AppConfig

func LoadEnvironment(ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	err = envconfig.Process(ctx, &Conf)
	if err != nil {
		log.Fatal("Error processing environment variables: ", err.Error())
	}
	log.Println("EnvConfig variables loaded")
}

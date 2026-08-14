package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB string
	ServerPort string
	GinEnv string
}


func Load() (*Config, error) {
	if err:=godotenv.Load();err!=nil{
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg := &Config{
		MongoURI:   os.Getenv("MONGO_URI"),
		MongoDB:    os.Getenv("MONGO_DB_NAME"),
		ServerPort: os.Getenv("PORT"),
		GinEnv:     os.Getenv("GIN_ENV"),
	}

	if cfg.MongoURI == "" || cfg.MongoDB == "" || cfg.ServerPort == "" || cfg.GinEnv == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}

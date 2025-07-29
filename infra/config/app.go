package config

import (
	"os"
	"time"
)

type AppConfig struct {
	ServerPort   string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MongoURI     string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		ServerPort:   ":8080",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MongoURI:     os.Getenv("MONGO_URI"),
	}
}

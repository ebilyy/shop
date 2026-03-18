package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN string `json:"db_dsn"`
}

func Load() Config {
	_ = godotenv.Load("../.env")
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN env var is required")
	}

	return Config{
		DBDSN: dsn,
	}
}
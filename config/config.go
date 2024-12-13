package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RajaOngkirAPIKey string
	RajaOngkirURL	 string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return &Config{
		RajaOngkirAPIKey: os.Getenv("RAJAONGKIR_APIKEY"),
		RajaOngkirURL : os.Getenv("RAJAONGKIR_URL"),
	}
}
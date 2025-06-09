package config

import (
	"log"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() Config {
	envPath := filepath.Join("configs", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("[config] Warning: Could not load .env file from %s: %v", envPath, err)
		log.Printf("[config] Falling back to system environment variables")
	} else {
		log.Printf("[config] Successfully loaded environment from %s", envPath)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("[config] Running on port %s", port)
	return Config{Port: port}
}

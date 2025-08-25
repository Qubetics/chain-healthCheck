package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetRequiredEnvVars(keys ...string) (map[string]string, error) {
	envs := make(map[string]string)
	for _, key := range keys {
		val := os.Getenv(key)
		if val == "" {
			return nil, fmt.Errorf("missing %s env var", key)
		}
		envs[key] = val
	}
	return envs, nil
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables.")
	}
}

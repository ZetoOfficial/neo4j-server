package config

import (
	"os"
)

type Config struct {
	Neo4jURI      string
	Neo4jUser     string
	Neo4jPassword string
	AuthToken     string
	HTTPPort      string
}

func LoadConfig() *Config {
	return &Config{
		Neo4jURI:      getEnv("NEO4J_URI", "bolt://localhost:7687"),
		Neo4jUser:     getEnv("NEO4J_USER", "neo4j"),
		Neo4jPassword: getEnv("NEO4J_PASSWORD", "password"),
		AuthToken:     getEnv("AUTH_TOKEN", "secret-token"),
		HTTPPort:      getEnv("HTTP_PORT", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

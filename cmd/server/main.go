package main

import (
	"log"

	"github.com/ZetoOfficial/neo4j-server/internal/config"
	"github.com/ZetoOfficial/neo4j-server/internal/delivery/http"
	"github.com/ZetoOfficial/neo4j-server/internal/repository"
	"github.com/ZetoOfficial/neo4j-server/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env not found: %v", err)
	}

	cfg := config.LoadConfig()

	repo, err := repository.NewNeo4jStorage(cfg.Neo4jURI, cfg.Neo4jUser, cfg.Neo4jPassword)
	if err != nil {
		log.Fatalf("Failed to connect to Neo4j: %v", err)
	}

	svc := service.NewGraphService(repo)
	log.Println("Успешно подключено к базе данных Neo4j")
	handler := http.NewHandler(svc)

	router := gin.Default()
	router.Use(cors.Default())
	handler.RegisterRoutes(router)

	if err := router.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

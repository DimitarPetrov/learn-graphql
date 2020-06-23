package main

import (
	"github.com/DimitarPetrov/learn-graphql/internal/handlers"
	"github.com/DimitarPetrov/learn-graphql/pkg/config"
	"github.com/DimitarPetrov/learn-graphql/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error creating config: %s", err)
	}

	server.Run(cfg.Server, []handlers.Handler{
		&handlers.PingHandler{},
		&handlers.GraphQLHandler{},
		&handlers.PlaygroundHandler{},
	})
}

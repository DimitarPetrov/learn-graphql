package server

import (
	"fmt"
	"github.com/DimitarPetrov/learn-graphql/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/vrischmann/envconfig"
	"log"
)

type Settings struct {
	Host string `envconfig:"default=localhost"`
	Port int    `envconfig:"default=7777"`
}

func (s *Settings) Validate() error {
	if len(s.Host) == 0 {
		return fmt.Errorf("server host missing")
	}
	if s.Port == 0 {
		return fmt.Errorf("server port missing")
	}
	return nil
}

func (s *Settings) Init() error {
	return envconfig.InitWithPrefix(&s, "SERVER")
}

func Run(settings Settings, handlers []handlers.Handler) {
	r := gin.Default()
	// Setup routes
	for _, handler := range handlers {
		route := handler.Route()
		r.Handle(route.Endpoint.Method, route.Endpoint.Path, route.HandlerFunc)
	}

	log.Printf("Running @ http://%s:%d", settings.Host, settings.Port)
	log.Fatalln(r.Run(fmt.Sprintf("%s:%d", settings.Host, settings.Port)))
}

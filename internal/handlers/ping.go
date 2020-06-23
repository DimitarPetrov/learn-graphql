package handlers

import (
	"github.com/DimitarPetrov/learn-graphql/pkg/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingHandler struct{}

// Ping is simple keep-alive/ping handler
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (*PingHandler) Route() Route {
	return Route{
		Endpoint: Endpoint{
			Method: http.MethodGet,
			Path:   routes.PingURL,
		},
		HandlerFunc: Ping,
	}
}

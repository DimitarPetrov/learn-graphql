package handlers

import (
	"github.com/gin-gonic/gin"
)

// Controller is an entity that wraps a set of HTTP Routes
type Handler interface {
	// Routes returns the set of routes for this controller
	Route() Route
}

// Route is a mapping between an Endpoint and a REST API Handler
type Route struct {
	// Endpoint is the combination of Path and HTTP Method for the specified route
	Endpoint Endpoint

	// Handler is the function that should handle incoming requests for this endpoint
	HandlerFunc gin.HandlerFunc
}

// Endpoint is a combination of a Path and an HTTP Method
type Endpoint struct {
	Method string
	Path   string
}

package handlers

import (
	gqhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DimitarPetrov/learn-graphql/internal/gql"
	"github.com/DimitarPetrov/learn-graphql/internal/gql/resolvers"
	"github.com/DimitarPetrov/learn-graphql/pkg/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GraphQLHandler struct {}

func (*GraphQLHandler) Route() Route {
	return Route{
		Endpoint: Endpoint{
			Method: http.MethodPost,
			Path: routes.GraphqlURL,
		},
		HandlerFunc: GraphqlHandlerFunc(),
	}
}

type PlaygroundHandler struct {}

func (*PlaygroundHandler) Route() Route {
	return Route{
		Endpoint: Endpoint{
			Method: http.MethodGet,
			Path: routes.PlaygroundURL,
		},
		HandlerFunc: PlaygroundHandlerFunc(routes.GraphqlURL),
	}
}

// GraphqlHandlerFunc defines the GQLGen GraphQL server handler
func GraphqlHandlerFunc() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := gqhandler.New(gql.NewExecutableSchema(c))
	h.AddTransport(transport.POST{})
	h.Use(extension.Introspection{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandlerFunc Defines the Playground handler to expose our playground
func PlaygroundHandlerFunc(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

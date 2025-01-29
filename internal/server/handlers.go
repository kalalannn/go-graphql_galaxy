package server

import (
	"encoding/json"
	"go-graphql_galaxy/internal/graphql/generated"
	"go-graphql_galaxy/internal/graphql/resolvers"
	"go-graphql_galaxy/pkg/utils"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func NewGraphQLHandler(config *utils.Server, resolver *resolvers.Resolver) *handler.Server {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	//! TODO redis cache
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	// Security for production
	if config.UseIntrospection {
		srv.Use(extension.Introspection{})
	}

	srv.Use(extension.FixedComplexityLimit(config.GQLComplexityLimit))

	srv.Use(NewDepthExtension(config.GQLDepthLimit))

	//! TODO add depth limit + rate limit + middleware

	// req -> resp(to_cache) + id ; id -> req(cached) -> resp
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func NewPlaygroundHandler(title, gqlPath string) http.Handler {
	return playground.Handler(title, gqlPath)
}

func NewPingHandler(pingMsg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": pingMsg}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})
}

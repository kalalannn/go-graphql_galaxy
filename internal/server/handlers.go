package server

import (
	"go-graphql_galaxy/internal/graphql/generated"
	"go-graphql_galaxy/internal/graphql/resolvers"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

const (
	GraphQLPlaygroundTitle = "GraphQL Playground"
	PlaygroundPath         = "/"
	GraphQLPath            = "/query"
)

type ServerService struct {
	Config *utils.Server
	DB     *gorm.DB
}

func NewServerService(config *utils.Server, db *gorm.DB) *ServerService {
	s := ServerService{
		Config: config,
		DB:     db,
	}
	s.initHandlers()
	return &s
}

func (s *ServerService) initHandlers() {
	http.Handle(GraphQLPath, s.NewGraphQLHandler())
	if s.Config.UsePlayground {
		http.Handle(PlaygroundPath, s.NewPlaygroundHandler(GraphQLPlaygroundTitle, GraphQLPath))
	}
}

func (s *ServerService) RunServer() error {
	host, port := s.Config.Host, s.Config.Port
	log.Info("Serving GraphQL on: http://%s:%s/", host, port)
	return http.ListenAndServe(host+":"+port, nil)
}

func (s *ServerService) NewGraphQLHandler() *handler.Server {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		DB: s.DB,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	//! TODO redis cache
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	// Security for production
	if s.Config.UseIntrospection {
		srv.Use(extension.Introspection{})
	}

	srv.Use(extension.FixedComplexityLimit(s.Config.GQLComplexityLimit))

	//! TODO add depth limit + rate limit + middleware

	// req -> resp(to_cache) + id ; id -> req(cached) -> resp
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func (s *ServerService) NewPlaygroundHandler(title, path string) http.HandlerFunc {
	return playground.Handler("GraphQL", "/query")
}

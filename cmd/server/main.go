package main

import (
	"fmt"
	"go-graphql_galaxy/graph"
	"go-graphql_galaxy/pkg/database"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	config := utils.MustLoadConfig()

	// Init logger
	log.Init(config.Env)

	// Connect to DB
	dsn := utils.DSN(&config.Database)
	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("failed to connect database: %v", err)
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL", "/query"))
	http.Handle("/query", srv)

	host, port := config.Server.Host, config.Server.Port
	log.Info("connect to http://%s:%s/ for GraphQL", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}

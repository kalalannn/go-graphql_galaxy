package server

import (
	"context"
	"fmt"
	"go-graphql_galaxy/internal/graphql/resolvers"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gorm.io/gorm"
)

const (
	GraphQLPlaygroundTitle = "GraphQL Playground"
	PlaygroundPath         = "/"
	GraphQLPath            = "/query"
	PingPath               = "/ping"
	PingMsg                = "pong"
)

const (
	ShutdownTimeout = 10 * time.Second
)

type ServerService struct {
	config *utils.Server
	db     *gorm.DB
	srv    *http.Server
}

func NewServerService(config *utils.Server, db *gorm.DB) *ServerService {
	s := ServerService{
		config: config,
		db:     db,
	}
	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: s.initRoutes(),
	}
	return &s
}

func (s *ServerService) initRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(PingPath, NewPingHandler(PingMsg))
	mux.Handle(GraphQLPath, NewGraphQLHandler(s.config, resolvers.NewResolver(s.db)))

	if s.config.UsePlayground {
		mux.Handle(PlaygroundPath, NewPlaygroundHandler(GraphQLPlaygroundTitle, GraphQLPath))
	}

	return mux
}

func (s *ServerService) RunServer(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	go func() {
		log.Info("Serving GraphQL on: %s", s.srv.Addr)

		if err := s.srv.ListenAndServe(); err != nil {
			log.Info("Stopped listening: %v", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	log.Warn("Shutting down server...")
	ctx, canel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer canel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Error("Shutdown with error: %v", err)
	}
	log.Info("Shutdown complete.")
}

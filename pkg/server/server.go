package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/cache"

	"github.com/shitcamp-unofficial/shitcamp/pkg/config"
	log "github.com/sirupsen/logrus"
)

type server struct {
	http.Server
}

func New() *server {
	return &server{}
}

func (s *server) Start() error {
	cfg := config.GetConfig()

	cache.Init(cfg.Cache.DefaultExpiryTime, cfg.Cache.DefaultCleanupInterval)

	s.Server.Addr = cfg.ServerAddress
	s.Server.Handler = newRouter()

	log.WithFields(log.Fields{"address": s.Server.Addr}).Info("starting_server")

	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("stopping_server")

	err := s.Shutdown(ctx)
	if err != nil {
		fmt.Println("error stopping API server:", err)
		log.WithError(err).Error("shutting_down_server")
		return err
	}

	return nil
}

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/kialkuz/task-manager/internal/api"
	"github.com/kialkuz/task-manager/internal/config"
	"github.com/kialkuz/task-manager/internal/infrastructure/env"
	"github.com/kialkuz/task-manager/pkg/logger"
)

func Handle(cfg config.ServerConfig) *http.Server {
	mux := api.Init(cfg.Routes...)

	return &http.Server{
		Addr:         ":" + env.GetEnv("TODO_PORT", ""),
		Handler:      mux,
		ErrorLog:     log.New(&logger.HttpErrorWriter{Logger: cfg.Logger}, "", log.LstdFlags),
		ReadTimeout:  time.Duration(5 * time.Second),
		WriteTimeout: time.Duration(10 * time.Second),
		IdleTimeout:  time.Duration(15 * time.Second),
	}
}

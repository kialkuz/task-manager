package api

import (
	"net/http"

	"github.com/kialkuz/task-manager/internal/infrastructure/env"
)

type RouteRegister interface {
	RegisterRoutes(mux *http.ServeMux)
}

func Init(routes ...RouteRegister) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir(env.GetEnv("WEB_DIR_PATH", ""))))
	mux.HandleFunc("POST /api/signin", signinHandler)

	for _, r := range routes {
		r.RegisterRoutes(mux)
	}

	return mux
}

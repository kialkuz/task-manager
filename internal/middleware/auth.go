package middleware

import (
	"log"
	"net/http"

	"github.com/kialkuz/task-manager/internal/infrastructure/env"
	jwtService "github.com/kialkuz/task-manager/pkg/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := env.GetEnv("AUTH_PASSWORD", "")
		if len(pass) > 0 {
			cookie, err := r.Cookie("token")
			if err != nil {
				log.Println("failed to sign jwt: %s\n" + err.Error())
				http.Error(w, "Ошибка аутентификации", http.StatusBadRequest)
				return
			}

			jwt := cookie.Value

			token, err := jwtService.CreateToken(pass)
			if err != nil {
				log.Println("failed to sign jwt: %s\n" + err.Error())
				http.Error(w, "Ошибка аутентификации", http.StatusBadRequest)
				return
			}

			if token != jwt {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
	"github.com/kialkuz/task-manager/internal/infrastructure/env"
	jwtService "github.com/kialkuz/task-manager/pkg/jwt"
)

type UserData struct {
	Password string `json:"password"`
}

type SuccessResponse struct {
	Token string `json:"token"`
}

const (
	tokenExpired = 8
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data UserData

	if err = json.Unmarshal(buf.Bytes(), &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Password != env.GetEnv("AUTH_PASSWORD", "") {
		httpService.WriteJson(w, responseDto.ErrorResponse{ErrorText: "Неверный пароль"}, http.StatusUnauthorized)
		return
	}

	token, err := jwtService.CreateToken(data.Password)
	if err != nil {
		log.Println("failed to sign jwt: %s\n" + err.Error())
		httpService.WriteJson(w, responseDto.ErrorResponse{ErrorText: "Неверный пароль"}, http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   token,
		Expires: time.Now().Add(tokenExpired * time.Hour),
	}
	http.SetCookie(w, cookie)

	httpService.WriteJsonOKResponse(w, SuccessResponse{Token: token})
}

package middleware

import (
	"context"
	"net/http"

	"github.com/fajarcandraaa/pizza_hub/internal/dto"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/jinzhu/gorm"
)

type MiddlewareStruct struct {
	db *gorm.DB
}

func MiddlewareStructFunc(db *gorm.DB) *MiddlewareStruct {
	return &MiddlewareStruct{
		db: db,
	}
}

func (m *MiddlewareStruct) BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, pass, m.db) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string, db *gorm.DB) bool {
	payload := dto.LoginRequest(username, password)
	r := repositories.NewChefRepositories(db)
	_, err := r.SignIng(context.Background(), payload)
	if err != nil {
		return false
	}

	return true
}

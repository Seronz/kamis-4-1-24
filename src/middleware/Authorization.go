package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	j "github.com/dgrijalva/jwt-go"
	jwt "github.com/seronz/api/src/utils/JWT"
)

func LoggerMiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/register" ||
			r.URL.Path == "/login" || r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}

		authorization_header := r.Header.Get("Authorization")

		if !strings.Contains(authorization_header, "") {
			log.Println("middleware : token contains invalid characters")
			res := map[string]string{"error middelware": "invalid token"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		token, err := jwt.JWTParser(authorization_header)
		if err != nil {
			log.Println(fmt.Errorf("middleware : %s", err))
			res := map[string]interface{}{"error middleware": err}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		claims, ok := token.Claims.(j.MapClaims)
		if !ok || !token.Valid {
			log.Println("middleware : invalid token")
			res := map[string]interface{}{"error middleware": errors.New("token invalid")}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		type contextKey string
		const contextWithKey contextKey = "user info"

		ctx := context.WithValue(r.Context(), contextWithKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

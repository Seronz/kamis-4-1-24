package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	j "github.com/dgrijalva/jwt-go"
	jwt "github.com/seronz/api/src/utils/JWT"
)

func LoggerMiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/register" ||
			r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorization_header := r.Header.Get("Authorization")
		if !strings.Contains(authorization_header, "") {
			fmt.Println("error disini")
			res := map[string]string{"error middelware": "invalid token"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		token, err := jwt.JWTParser(authorization_header)
		fmt.Println("ini tokennya", token)
		if err != nil {
			fmt.Println("error disini 2")
			res := map[string]interface{}{"error middleware": err}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		claims, ok := token.Claims.(j.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("error disini 3")
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

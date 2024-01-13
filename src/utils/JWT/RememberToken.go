package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	jwt.StandardClaims
	ID         string `json:"id"`
	FIRSTNAME  string `json:"firstname"`
	LASTNAME   string `json:"lastname"`
	EMAIL      string `json:"email"`
	USERROLE   string `json:"userrole"`
	REMEMBERME bool   `json:"remember_me"`
}

type Params struct {
	W          http.ResponseWriter
	ID         string
	Firstname  string
	Lastname   string
	Email      string
	Userrole   string
	RememberMe bool
}

type JWTConfig struct {
	ApplicationName       string
	LogExpirationDuration time.Duration
	JWTSigningMethod      *jwt.SigningMethodHMAC
	JWTSignatureKey       []byte
}

func (p *Params) loadEnvJWT() (*JWTConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var ExpDuration time.Duration
	if p.Userrole == "" {
		ExpDuration = time.Duration(2) * time.Hour
	} else if p.Userrole == "user" || p.Userrole == "admin" {
		ExpDuration = time.Duration(48) * time.Hour
	}

	return &JWTConfig{
		ApplicationName:       os.Getenv("APPLICATION_NAME"),
		JWTSigningMethod:      jwt.SigningMethodHS256,
		LogExpirationDuration: ExpDuration,
		JWTSignatureKey:       []byte(os.Getenv("SIGNATURE_KEY")),
	}, nil
}

// func loadJWTConfig(params Params) (*JWTConfig, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var ExpDuration time.Duration

// 	if params.Userrole == "" {
// 		ExpDuration = time.Duration(2) * time.Hour
// 	} else if params.Userrole == "user" || params.Userrole == "admin" {
// 		ExpDuration = time.Duration(48) * time.Hour
// 	}

// 	return &JWTConfig{
// 		ApplicationName:       os.Getenv("APPLICATION_NAME"),
// 		JWTSigningMethod:      *jwt.SigningMethodHS256,
// 		LogExpirationDuration: ExpDuration,
// 		JWTSignatureKey:       []byte(os.Getenv("SIGNATURE_KEY")),
// 	}, nil

// }

func generateRememberToken(param Params) (MyClaims, error) {
	var p Params

	p.RememberMe = param.RememberMe
	config, err := p.loadEnvJWT()
	if err != nil {
		return MyClaims{}, err
	}

	remember := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.ApplicationName,
			ExpiresAt: time.Now().Add(time.Duration(168) * time.Hour).Unix(),
			Subject:   param.Email,
		},
		ID:         param.ID,
		FIRSTNAME:  param.Firstname,
		LASTNAME:   param.Lastname,
		EMAIL:      param.Email,
		USERROLE:   param.Userrole,
		REMEMBERME: param.RememberMe,
	}
	return remember, nil
}

func generateClaims(param Params) (MyClaims, error) {
	var p Params

	p.Userrole = param.Userrole

	config, err := p.loadEnvJWT()
	if err != nil {
		return MyClaims{}, err
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.ApplicationName,
			ExpiresAt: time.Now().Add(config.LogExpirationDuration).Unix(),
			Subject:   param.Email,
		},
		FIRSTNAME: param.Firstname,
		LASTNAME:  param.Lastname,
		EMAIL:     param.Email,
		USERROLE:  param.Userrole,
	}

	return claims, nil
}

func CreateRememberToken(param Params) (string, error) {
	remember, err := generateRememberToken(param)
	if err != nil {
		return "", err
	}

	var p Params
	p.RememberMe = param.RememberMe
	config, err := p.loadEnvJWT()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		config.JWTSigningMethod,
		remember,
	)
	rememberToken, err := token.SignedString(config.JWTSignatureKey)
	if err != nil {
		return "", err
	}
	return rememberToken, nil
}

func CreateToken(params Params) (string, error) {
	claims, err := generateClaims(params)
	if err != nil {
		return "", err
	}
	var p Params

	p.Userrole = params.Userrole
	config, err := p.loadEnvJWT()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		config.JWTSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(config.JWTSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JWTParser(header string) (*jwt.Token, error) {
	var p Params
	res, err := p.loadEnvJWT()
	if err != nil {
		return nil, err
	}
	fmt.Println("ini signing method:\n", res.JWTSigningMethod)
	fmt.Println("ini application name:", res.ApplicationName)
	token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != res.JWTSigningMethod {
			fmt.Println("errornya disini kak")
			return nil, errors.New("method invalid")
		}
		return res.JWTSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

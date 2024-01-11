package jwt

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	jwt.StandardClaims
	FIRSTNAME string `json:"firstname"`
	LASTNAME  string `json:"lastname"`
	EMAIL     string `json:"email"`
	USERROLE  string `json:"userrole"`
}

type Params struct {
	W         http.ResponseWriter
	Firstname string
	Lastname  string
	Email     string
	Userrole  string
}

type JWTConfig struct {
	ApplicationName       string
	LogExpirationDuration time.Duration
	JWTSigningMethod      jwt.SigningMethodHMAC
	JWTSignatureKey       []byte
}

func loadJWTConfig(params Params) (*JWTConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var ExpDuration time.Duration

	if params.Userrole == "guest" {
		ExpDuration = time.Duration(2) * time.Hour
	} else {
		ExpDuration = time.Duration(48) * time.Hour
	}

	return &JWTConfig{
		ApplicationName:       os.Getenv("APPLICATION_NAME"),
		JWTSigningMethod:      *jwt.SigningMethodHS256,
		LogExpirationDuration: ExpDuration,
		JWTSignatureKey:       []byte(os.Getenv("SIGNATURE_KEY")),
	}, nil

}

func generateClaims(param Params) (MyClaims, error) {
	config, err := loadJWTConfig(param)
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

func CreateToken(params Params) (string, error) {
	claims, err := generateClaims(params)
	if err != nil {
		return "", err
	}

	config, err := loadJWTConfig(params)
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(
		&config.JWTSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(config.JWTSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

package otp

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pquerna/otp/totp"
)

func GenerateOTP(email string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	APPLICATION_NAME := os.Getenv("APPLICATION_NAME")

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      APPLICATION_NAME,
		AccountName: email,
		Period:      7200,
	})

	if err != nil {
		return "", err
	}
	// url := key.URL()

	// fmt.Println(url)

	otpCode, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", err
	}
	return otpCode, nil
}

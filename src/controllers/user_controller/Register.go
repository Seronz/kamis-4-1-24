package usercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/seronz/api/config"
	users "github.com/seronz/api/src/models/Users"
	response "github.com/seronz/api/src/utils/Response"
)

func User(w http.ResponseWriter, r *http.Request) {

	var User users.User
	err := json.NewDecoder(r.Body).Decode(&User)

	log.Printf("register : %s %s", User.FirstName, User.LastName)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: "gagal mendapatkan body",
		}
		response.ResponseFailed(res)
	}

	data, token, otp, err := users.UserRegister(config.Mongo, User)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: "gagal mendapatkan body",
		}
		response.ResponseFailed(res)
	}

	w.Header().Set("authorization", token)
	w.Header().Set("Access-Control-Expose-Headers", "authorization")
	res := response.Response{
		W:        w,
		Data:     data,
		Messages: otp,
	}
	response.ResponseSuccess(res)
}

func ActivateAccount(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(fmt.Errorf("activate account :  %s", err))
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
		return
	}
	err = users.ActivationAccount(config.DB, config.Mongo, user.OtpCode)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
		return
	}

	res := response.Response{
		W:        w,
		Messages: "Success",
	}
	response.ResponseSuccess(res)
}

func GenerateOTP(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")

	otp, err := users.RegenerateOTP(config.Mongo, tokenStr)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
		return
	}

	res := response.Response{
		W:        w,
		Data:     otp,
		Messages: "Success",
	}
	response.ResponseSuccess(res)
}

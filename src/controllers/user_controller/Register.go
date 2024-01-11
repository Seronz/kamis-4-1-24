package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/seronz/api/config"
	users "github.com/seronz/api/src/models/Users"
	response "github.com/seronz/api/src/utils/Response"
)

func User(w http.ResponseWriter, r *http.Request) {
	var User users.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: "gagal mendapatkan body",
		}
		response.ResponseFailed(res)
	}

	data, err := users.UserRegister(config.DB, config.Mongo, User)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: "gagal mendapatkan body",
		}
		response.ResponseFailed(res)
	}
	res := response.Response{
		W:        w,
		Data:     data,
		Messages: "Success",
	}
	response.ResponseSuccess(res)
}

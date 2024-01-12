package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/seronz/api/config"
	users "github.com/seronz/api/src/models/Users"
	response "github.com/seronz/api/src/utils/Response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
	}

	result, err := users.Login(config.DB, w, user)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
	}

	w.Header().Set("Authorization", result)
	w.Header().Set("Access-Control-Expose-Headers", "authorization")
}

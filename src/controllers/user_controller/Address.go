package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/seronz/api/config"
	users "github.com/seronz/api/src/models/Users"
	response "github.com/seronz/api/src/utils/Response"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address users.Address

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)

	}

	user_token := r.Header.Get("Authorization")

	_, err = users.CreateAddress(config.DB, address, user_token)
	if err != nil {
		res := response.Response{
			W:        w,
			Err:      err,
			Messages: err.Error(),
		}
		response.ResponseFailed(res)
	}

}

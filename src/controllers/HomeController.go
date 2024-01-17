package controllers

import (
	"fmt"
	"net/http"

	response "github.com/seronz/api/src/utils/Response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	res := response.Response{
		W: w,

		Messages: "hallo bro!!!",
	}
	response.ResponseFailed(res)
	fmt.Println("test")
}

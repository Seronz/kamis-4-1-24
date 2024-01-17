package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	W        http.ResponseWriter
	Status   int
	Data     interface{}
	Messages string
	Err      error
}

func writeJsonResponse(response Response) {
	response.W.Header().Set("Content-Type", "application/json")
	response.W.WriteHeader(response.Status)

	r := map[string]interface{}{
		"response": response.Data,
		"messages": response.Messages,
	}

	json.NewEncoder(response.W).Encode(r)
}

func ResponseSuccess(response Response) {
	r := Response{
		W:        response.W,
		Status:   http.StatusOK,
		Data:     response.Data,
		Messages: response.Messages,
	}
	writeJsonResponse(r)
}

func ResponseFailed(response Response) {
	r := Response{
		W:        response.W,
		Status:   http.StatusBadRequest,
		Data:     response.Err,
		Messages: response.Messages,
	}

	writeJsonResponse(r)
}

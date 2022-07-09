package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ferjoaguilar/testing-example-golang/models"
)

func newResponse(message string, data any) *models.Response {
	return &models.Response{Message: message, Data: data}
}

func ResponseError(res http.ResponseWriter, statusCode int, message string, data any) error {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	httpResponse := newResponse(message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}

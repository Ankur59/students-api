package response

import (
	"encoding/json"
	"net/http"
)

type errorStructure struct {
	Status  int
	Message string
	Data    interface{}
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(status int, message string, data interface{}) errorStructure {
	return errorStructure{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

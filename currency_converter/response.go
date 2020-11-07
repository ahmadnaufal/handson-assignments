package currencyconverter

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string
}

func OK(w http.ResponseWriter, data interface{}) error {
	response, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return nil
}

func Created(w http.ResponseWriter, data interface{}) error {
	response, _ := json.Marshal(data)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	return nil
}

func Error(w http.ResponseWriter, err error) error {
	response, _ := json.Marshal(errorResponse{err.Error()})
	status := http.StatusInternalServerError

	w.WriteHeader(status)
	w.Write(response)

	return err
}

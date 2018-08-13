package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// ApplicationException - application error structure
type ApplicationException struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

// InternalServiceError - generic error with application
var InternalServiceError = ApplicationException{
	Code:       "InternalServiceError",
	Message:    "Something went wrong!",
	HTTPStatus: http.StatusInternalServerError,
}

func handleError(e error, appError ApplicationException, w http.ResponseWriter) {
	log.Println(e)
	log.Println(appError.Code)

	ae, err := json.Marshal(appError)
	if err != nil {
		log.Fatal(e)
	}

	w.WriteHeader(appError.HTTPStatus)
	w.Write(ae)
}

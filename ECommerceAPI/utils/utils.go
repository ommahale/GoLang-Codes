package utils

import (
	"log"
	"net/http"
)

func HandleError(err error, mssg string) {
	if err != nil {
		log.Fatal(mssg, " ", err.Error())
	}
}
func HandleHttpError(err error, mssg string, w *http.ResponseWriter) {
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}

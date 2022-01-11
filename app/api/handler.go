package api

import (
	"log"

	"github.com/adeniyistephen/cloudkite/app/business"
	"github.com/gorilla/mux"
)

func Handle(log *log.Logger) *mux.Router {
	r := mux.NewRouter()

	// Register passenger endpoints.
	mg := MessageGroup{
		message: business.New(log),
	}

	r.HandleFunc("/hello", mg.RespondHello).Methods("GET")
	r.HandleFunc("/vowel-service", mg.MessageCreate).Methods("POST")

	return r
}

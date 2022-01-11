package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adeniyistephen/cloudkite/app/business"
	"github.com/pkg/errors"
)

// MessageGroup is a group of message related operations.
type MessageGroup struct {
	message business.NewCloudKiteService
}

// MessageCreate creates a new reverse vowel message.
func (mg MessageGroup) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	decoder := json.NewDecoder(r.Body)
	var m business.Message
	if err := decoder.Decode(&m); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Create a new reverse message.
	Rmsg, err := mg.message.ReverseVowels(m)
	if err != nil {
		log.Panic(errors.Wrap(err, "failed to reverse vowels"))
	}

	respondWithJSON(w, http.StatusCreated, Rmsg)
}

func (mg MessageGroup) RespondHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	// Respond to hello handler.
	res, err := mg.message.Hello()
	if err != nil {
		log.Panic(errors.Wrap(err, "failed to get hello response"))
	}

	respondWithJSON(w, http.StatusCreated, res)
}

// respondWithError returns an error message.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON returns a JSON response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

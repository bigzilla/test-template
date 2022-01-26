package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var (
	// the author of package schema recommend to put the decoder instance as a package global
	// see the reason behind it here: https://github.com/gorilla/schema#example
	decoder = schema.NewDecoder()

	// use a single instance of Validate, it caches struct info
	validate = validator.New()
)

// decodeSchema decode schema from url query into v and validate the field, v must be pointer to a struct
func decodeSchema(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := decoder.Decode(v, r.URL.Query()); err != nil {
		return fmt.Errorf("failed to decode url query: %v", err)
	}
	return validate.Struct(v)
}

// decodeJSON decode json from request body into v and validate the field, v must be pointer to a struct
func decodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode request body: %v", err)
	}
	return validate.Struct(v)
}

// respondJSON responds to the request with specific HTTP code and JSON data
// log the error if data is error type and the JSON response is {"error": data}, otherwise {"data": data}
func respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(code)

	if err, ok := data.(error); ok {
		data = map[string]interface{}{"errors": err.Error()}
	} else {
		data = map[string]interface{}{"data": data}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloWorldResponseStruct struct {
	MsgID   int    `json:"msgID"`
	Message string `json:"message"`
}

type helloWorldRequestStruct struct {
	Name string `json:"name"`
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("In Validation Handler")
	h.next.ServeHTTP(w, r)
}

type mainHandler struct{}

func newMainHandler() http.Handler {
	return mainHandler{}
}

func (h mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("In Main Handler")
	var request helloWorldRequestStruct
	var response helloWorldResponseStruct

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)

	w.Header().Set("Content-Type", "application/json")
	response.MsgID = 1
	response.Message = "Hello there " + request.Name
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Println("Error Encoding JSON : %v", err)
	}
}

func main() {
	log.Println("Starting the Server on port 8080...")

	handler := newValidationHandler(newMainHandler())

	http.Handle("/hello", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

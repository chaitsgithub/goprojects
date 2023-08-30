package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	log.Println("Starting Server on 8080 port...")

	// http.HandleFunc("/hello", MarshallHandler)
	http.HandleFunc("/hello", EncoderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MarshallHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequestStruct
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	err = json.Unmarshal(reqData, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	var response helloWorldResponseStruct
	response.MsgID = 1
	response.Message = "Hello There " + request.Name
	// data, err := json.Marshal(response)
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println("Error Marshalling JSON : %v", err)
	}
	fmt.Fprintf(w, string(data))
}

func EncoderHandler(w http.ResponseWriter, r *http.Request) {
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

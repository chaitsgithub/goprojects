package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.google.com/", nil)
	if err != nil {
		log.Println(err)
	}

	ctx, cancelCtx := context.WithTimeout(req.Context(), 250*time.Millisecond)
	defer cancelCtx()
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Successful Response received. Status : %v", resp.Status)
}

package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	log.Println("Starting elasticsearch with go...")

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating elasticsearch client: %s", err)
	}
	log.Println(elasticsearch.Version)

	resp, err := es.Info()
	if err != nil {
		log.Fatalf("Error printing elasticsearch info: %s", err)
	}
	defer resp.Body.Close()
	log.Println(resp)
}

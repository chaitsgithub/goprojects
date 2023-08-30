package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var es, _ = elasticsearch.NewDefaultClient()

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func ReadText(reader *bufio.Scanner, prompt string) string {
	fmt.Print(prompt + ": ")
	reader.Scan()
	return reader.Text()
}

func LoadData() {
	var spacecrafts []map[string]interface{}
	pageNumber := 0
	for {
		response, _ := http.Get("http://stapi.co/api/v1/rest/spacecraft/search?pageSize=100&pageNumber=" + strconv.Itoa(pageNumber))
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		var result map[string]interface{}
		json.Unmarshal(body, &result)

		page := result["page"].(map[string]interface{})
		totalPages := int(page["totalPages"].(float64))

		crafts := result["spacecrafts"].([]interface{})
		log.Printf("Processing %d / %d \n", pageNumber+1, totalPages)

		for _, craftInterface := range crafts {
			craft := craftInterface.(map[string]interface{})
			spacecrafts = append(spacecrafts, craft)
		}
		pageNumber++
		if pageNumber >= totalPages {
			break
		}
	}

	log.Println("Loading crafts into ElasticSearch...")
	for _, data := range spacecrafts {
		uid := data["uid"].(string)
		jsonString, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Error marshalling data : %s", err)
		}
		request := esapi.IndexRequest{Index: "stsc", DocumentID: uid, Body: strings.NewReader(string(jsonString))}
		_, err = request.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error indexing data : %s", err)
		}
	}
	log.Println("Loading crafts to ElasticSearch completed!")
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load Spacecraft")
		fmt.Println("2) Get Spacecraft")
		option := ReadText(reader, "Enter option")
		if option == "0" {
			Exit()
		} else if option == "1" {
			LoadData()
		} else {
			fmt.Println("Invalid option")
		}
	}
}

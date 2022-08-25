package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// basicData(client)
	structData(client)

}

func structData(client *redis.Client) {
	auth1 := Author{Name: "Chaitanya", Age: 40}
	jsonout, err := json.Marshal(auth1)
	if err != nil {
		fmt.Printf("Error Marshalling Data : %v", err)
	}

	err = client.Set("ID1", jsonout, 0).Err()
	if err != nil {
		fmt.Printf("Error setting key in Redis : %v", err)
	}

	val, err := client.Get("ID1").Result()
	if err != nil {
		fmt.Printf("Error getting key in Redis : %v", err)
	}
	fmt.Println(val)
}

func basicData(client *redis.Client) {
	err := client.Set("Chaits1", "Awesome1", 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("Chaits1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

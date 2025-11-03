package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/**
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
**/

type apiResponse struct {
	UserId    int  `json:"userId"`
	ID        int  `json:"id"`
	Completed bool `json:"completed"`
}

func main() {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Could not create the request due to: ", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not perform the request due to: ", err.Error())
	}

	defer resp.Body.Close()

	// Decode the JSON response
	var data apiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Printf("Response Body (Golang struct): %+v \n", data)
}

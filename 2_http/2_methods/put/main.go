package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// URL for the PUT
const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	payload := []byte(`{"userId": 1, "id": 1, "title": "sunt aut facere", "body": "quia et suscipit"}`)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating Put request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Put request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Put Response code:", resp.StatusCode)
	fmt.Println("Put Response body:", string(body))
}

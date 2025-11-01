package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Specify the URL for the Delete request
const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	// Make the Delete request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error creating Delete request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Put request:", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Code: ", resp.StatusCode)
	fmt.Println("Delete Response:", string(body))
}

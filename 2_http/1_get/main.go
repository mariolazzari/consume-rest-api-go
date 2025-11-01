package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating GET request", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error during GET request", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading GET response", err)
		return
	}

	fmt.Println("Response code:", resp.StatusCode)
	fmt.Println("Response body:", string(body))

}

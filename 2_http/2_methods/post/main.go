package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// URL for the POST
const url = "https://jsonplaceholder.typicode.com/todos"

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	payload := []byte(`{"title": "foo", "body": "bar", "userId": 1000}`)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating client", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making client", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body", err)
		return
	}

	fmt.Println("Response code", res.StatusCode)
	fmt.Println("Response body", string(body))
}

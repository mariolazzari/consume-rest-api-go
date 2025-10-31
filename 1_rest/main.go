package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error reading resource", err)
		return
	}

	defer resp.Body.Close()

	// read body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return
	}

	fmt.Printf("Response status code: %d\n", resp.StatusCode)
	fmt.Println("Response body:", string(body))
}

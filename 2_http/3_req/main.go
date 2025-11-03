package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const requestURL = "https://jsonplaceholder.typicode.com/posts"

type TodoClient struct {
	httpClient http.Client
}

func main() {
	todoClient := TodoClient{
		httpClient: http.Client{
			Timeout: 3 * time.Second,
		},
	}

	todoClient.sendJSONRequest()
	todoClient.sendFormData()
}

type CreatePostRequestBody struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (c TodoClient) sendJSONRequest() {

	// Create a request body
	createRequestBody := CreatePostRequestBody{UserId: 1, ID: 2, Title: "Good Request", Body: "Obligatory Bag Of Words"}
	createRequestJSONPayload, err := json.Marshal(createRequestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a POST request with JSON body
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, requestURL, bytes.NewBuffer(createRequestJSONPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error occured while performing request: ", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Println("Response Body:", string(body))

}

func (c TodoClient) sendFormData() {
	// Create form data
	formData := url.Values{
		"firstName": {"Roger"},
		"lastName":  {"Smith"},
	}

	// Create a POST request with form data
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, requestURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request...
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error occured while performing request: ", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Println("Response Body:", string(body))
}

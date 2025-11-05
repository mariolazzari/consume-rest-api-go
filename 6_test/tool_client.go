package chapter6

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Client represents the HTTP client interface.
type Client struct {
	HTTPClient HTTPClient
}

// HTTPClient is an interface representing the methods of http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// FetchData fetches data from an external API using the HTTP client.
func (c *Client) FetchTodo() (Todo, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		return Todo{}, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return Todo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Todo{}, errors.New("unexpected status code")
	}

	// Decode the JSON response
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Todo{}, err
	}

	return todo, nil
}

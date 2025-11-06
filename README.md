# Consuming REST Api with Go

## RESTfull Api

### Introduction to REST Api

Representational state transfer: set of architectural principles for designing network applications

#### Request components

- HTTP methods
- URI (uniform resource identifier)
- HTTP version
- Headers
- Request body (payload)
- Query parameters

#### REST constraints

- Client / Server
- Uniform interface
- Cache ability
- Layered system

#### Response types

- POST: creates new resource
- GET: get existing resource
- PUT: update full existing existing resource
- PATCH: partially update existing resource
- DELETE: delete existing resource

#### Server operations

- Success:
  - 200: success
  - 201: new redource created
  - 202: request accepted
- Messages:
  - 300
  - 301: resource moved
- Client errors
  - 400: bad request
  - 401: Invalid credentials
  - 403: no permissions
  - 404: resource not found
  - 405: method not allowed
- Server error
  - 500: internal server error

### Go native http client

[Doc](https://pkg.go.dev/net/http)

```go
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
```

#### Consideration for http client

- Error handling
- Authentication
- Other HTTP methods
- Timeouts

## Making htttp requests

### GET request

```go
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

	defer res.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading GET response", err)
		return
	}

	fmt.Println("Response code:", resp.StatusCode)
	fmt.Println("Response body:", string(body))

}
```

### Making POST, PUT, PATCH and DELETE

#### POST request

```go
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
```

#### PUT request

```go
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
```

#### PATCH request

```go
package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// URL for the PATCH
const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	payload := []byte(`{"userId": 1, "id": 1, "title": "sunt aut facere", "body": "quia et suscipit"}`)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating Patch request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Patch request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Patch Response code:", resp.StatusCode)
	fmt.Println("Patch Response body:", string(body))
}
```

#### DELETE reques

```go
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
```

### Handling request headers and bodies

```go
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
```

## Handling API responses

### Parsing JSON responses

Unmarshalling: convert JSON object to Go struct

```go
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
```

### Error handling

- Network errors
- Response error codes
- JSON parsing errors

#### Tips

- Logs error for debugging
- Explicit error handling
- Error wrapping
- Retry for transient error
- Use context for cancellation
- Test all possible errors

```go
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
```

## Authentication and authorization

### Basic authentication

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Get credentials from environment variables
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// Use the credentials
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
```

## Advanced topics

### Concurrency in API requests

```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var listOfEndpoints = [5]string{
	"https://jsonplaceholder.typicode.com/posts",
	"https://jsonplaceholder.typicode.com/albums",
	"https://jsonplaceholder.typicode.com/users",
	"https://jsonplaceholder.typicode.com/photos",
	"https://jsonplaceholder.typicode.com/comments",
}

func main() {
	fetchDataSequentially()
	fetchDataConcurrently()
	fetchDataConcurrentlyWithBufferedChannels()
}

func fetchData(endpoint string) ([]byte, error) {

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Println("Could not create the request due to: ", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not perform the request due to: ", err.Error())
	}

	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	return body, nil
}

func fetchDataSequentially() {
	defer timeTrack(time.Now(), "Sequential Data Fetch Operation")
	for _, endpoint := range listOfEndpoints {
		_, err := fetchData(endpoint)
		if err != nil {
			log.Printf("Error fetching data from %s: %v", endpoint, err)
			continue
		}
	}
}

func fetchDataConcurrently() {
	defer timeTrack(time.Now(), "Concurrent Data Fetching Operation With WaitGroup")
	var wg sync.WaitGroup
	for _, endpoint := range listOfEndpoints {
		wg.Add(1)
		go func(ep string) {
			defer wg.Done()
			_, err := fetchData(ep)
			if err != nil {
				log.Printf("Error fetching data from %s: %v \n", ep, err)
				return
			}
			// Process the data
			// log.Printf("Gotten Data from %s: %+v \n", ep, string(data))
		}(endpoint)
	}
	wg.Wait()
}

func fetchDataConcurrentlyWithBufferedChannels() {
	defer timeTrack(time.Now(), "Concurrent Data Fetching Operation With Buffered Channels")
	type Result struct {
		Data []byte
		Err  error
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(listOfEndpoints))
	for _, endpoint := range listOfEndpoints {
		wg.Add(1)
		go func(ep string) {
			defer wg.Done()
			data, err := fetchData(ep)
			result := Result{Data: data, Err: err}
			// Send the result to the channel
			results <- result
		}(endpoint)
	}
	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()
	// Process results from the channel
	for result := range results {
		if result.Err != nil {
			log.Printf("Error fetching data: %v", result.Err)
			continue
		}
		// Process the data
		// log.Printf("Gotten Data from %+v \n", string(result.Data))
	}
}

func timeTrack(start time.Time, operationName string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", operationName, elapsed)
}
```

### Rate limiting

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	maxAttempts = 5
	endpoint    = "https://jsonplaceholder.typicode.com/unknwon"
)

func main() {
	fetchDataWithExponentialBackoff()
}

func fetchDataWithExponentialBackoff() {
	client := &http.Client{}

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := client.Get(endpoint)
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			return
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			fmt.Printf("Rate limited. Retrying in %d seconds... \n", attempt)
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Request Not successful. Retrying in %d seconds... \n", attempt)
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}
		// Process the response
		// ...
		resp.Body.Close()
		break
	}
}

func fetchDataWithRateLimiter() {
	limiter := rate.NewLimiter(rate.Limit(10), 1) // 10 requests per second

	for i := 0; i < maxAttempts; i++ {
		if err := limiter.Wait(context.Background()); err != nil {
			fmt.Printf("Rate limited. Retrying in %v... \n", err)
			continue
		}

		resp, err := http.Get(endpoint)
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			return
		}
		// Process the response

		resp.Body.Close()
	}
}
```

### Optimizations

- Connection reuse and polling
- Caching responses
- Using HTTP keep-alive

## Testing

[httptest](https://pkg.go.dev/net/http/httptest)

### Unit testing

```go
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
```

```go
package chapter6

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchTodo(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the response for the test case
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}`))
	}))
	defer server.Close()

	// Make a "real" GET request to our mock http server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response status code to be %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the JSON response
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		t.Fatal(err)
	}

	// Check the fetched todo
	expectedTodo := Todo{UserId: 1, ID: 1, Title: "delectus aut autem", Completed: false}
	if todo != expectedTodo {
		t.Errorf("FetchTodo() returned unexpected todo. Expected: %+v, Got: %+v", expectedTodo, todo)
	}
}
```

### Mocking RST API

```go
package chapter6

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, errors.New("mock DoFunc not implemented")
}

func TestFetchTodoWithMockClient(t *testing.T) {
	// Create a mock client
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Define the response for the test case
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}`)),
			}, nil
		},
	}

	// Use the mock client in your client code
	client := &Client{HTTPClient: mockClient}

	// Call the FetchTodo method
	todo, err := client.FetchTodo()
	if err != nil {
		t.Errorf("FetchTodo() returned error: %v", err)
	}

	// Check the fetched todo
	expectedTodo := Todo{UserId: 1, ID: 1, Title: "delectus aut autem", Completed: false}
	if todo != expectedTodo {
		t.Errorf("FetchTodo() returned unexpected todo. Expected: %+v, Got: %+v", expectedTodo, todo)
	}
}
```

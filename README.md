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

	defer resp.Body.Close()

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

```go

```

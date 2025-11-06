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

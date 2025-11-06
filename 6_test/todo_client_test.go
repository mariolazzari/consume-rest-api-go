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

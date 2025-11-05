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

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

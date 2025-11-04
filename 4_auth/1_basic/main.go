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

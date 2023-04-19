package main

import (
	"fmt"
	"log"

	"github.com/yourusername/gojq/filter"
	"github.com/yourusername/gojq/jq"
)

func main() {
	// Example JSON input
	input := `{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "New York",
			"state": "NY"
		}
	}`

	// Create a filter
	f, err := filter.New(".address.city")
	if err != nil {
		log.Fatalf("Error creating filter: %v", err)
	}

	// Use ExecuteString to apply the filter to the input JSON
	result, err := jq.ExecuteString(f, input)
	if err != nil {
		log.Fatalf("Error executing jq filter: %v", err)
	}

	fmt.Printf("Result: %s\n", result)
}

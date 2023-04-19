package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yourusername/gojq/jq"
)

func main() {
	// Parse command-line arguments
	filter := flag.String("filter", "", "jq filter to apply")
	file := flag.String("file", "", "Path to JSON input file")
	flag.Parse()

	if *filter == "" {
		log.Fatalf("Filter is required")
	}

	if *file == "" {
		log.Fatalf("File is required")
	}

	// Read input JSON file
	input, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Use Execute to apply the filter to the input JSON
	result, err := jq.Execute(*filter, input)
	if err != nil {
		log.Fatalf("Error executing jq filter: %v", err)
	}

	fmt.Printf("Result: %s\n", result)
}

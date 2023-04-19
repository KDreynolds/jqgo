package filter

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// New creates a new jq filter.
func New(filter string) (string, error) {
	err := Validate(filter)
	if err != nil {
		return "", err
	}
	return filter, nil
}

// Validate checks if a given filter string is a valid jq filter.
func Validate(filter string) error {
	testInput := "{}"
	cmd := exec.Command("jq", filter)
	cmd.Stdin = strings.NewReader(testInput)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Invalid jq filter: %s. Error: %v", filter, stderr.String())
	}
	return nil
}

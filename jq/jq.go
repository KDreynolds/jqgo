package jq

import (
	"bytes"
	"errors"
	"os/exec"
)

// Execute runs a jq filter on a given JSON input and returns the result.
func Execute(filter string, input []byte) ([]byte, error) {
	cmd := exec.Command("jq", filter)
	cmd.Stdin = bytes.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, errors.New(stderr.String())
	}

	return stdout.Bytes(), nil
}

// ExecuteString runs a jq filter on a given JSON input string and returns the result as a string.
func ExecuteString(filter string, input string) (string, error) {
	result, err := Execute(filter, []byte(input))
	if err != nil {
		return "", err
	}

	return string(result), nil
}

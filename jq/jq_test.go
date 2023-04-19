package jq

import (
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	testCases := []struct {
		name         string
		filter       string
		input        []byte
		expected     []byte
		expectError  bool
	}{
		{
			name:         "Simple filter",
			filter:       ".name",
			input:        []byte(`{"name": "John Doe"}`),
			expected:     []byte(`"John Doe"`),
			expectError:  false,
		},
		{
			name:         "Invalid filter",
			filter:       "invalid",
			input:        []byte(`{"name": "John Doe"}`),
			expectError:  true,
		},
		{
			name:         "Invalid JSON",
			filter:       ".name",
			input:        []byte(`invalid`),
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Execute(tc.filter, tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !bytes.Equal(result, tc.expected) {
					t.Errorf("Expected result %q, got %q", tc.expected, result)
				}
			}
		})
	}
}

func TestExecuteString(t *testing.T) {
	testCases := []struct {
		name         string
		filter       string
		input        string
		expected     string
		expectError  bool
	}{
		// Add test cases similar to TestExecute, but with string input and expected values.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ExecuteString(tc.filter, tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if result != tc.expected {
					t.Errorf("Expected result %q, got %q", tc.expected, result)
				}
			}
		})
	}
}

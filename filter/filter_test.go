package filter

import (
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name        string
		filter      string
		expectError bool
	}{
		{
			name:        "Valid filter",
			filter:      ".name",
			expectError: false,
		},
		{
			name:        "Invalid filter",
			filter:      "invalid",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := New(tc.filter)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		name        string
		filter      string
		expectError bool
	}{
		{
			name:        "Valid filter",
			filter:      ".name",
			expectError: false,
		},
		{
			name:        "Invalid filter",
			filter:      "invalid",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.filter)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

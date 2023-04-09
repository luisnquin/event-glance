package validate_test

import (
	"testing"

	"github.com/luisnquin/event-glance/internal/validate"
)

func TestValidateAWSRegion(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Region", "us-west-2", true},
		{"Invalid Region", "invalid-region", false},
		{"Valid GovCloud Region", "us-gov-west-1", true},
		{"Invalid GovCloud Region", "us-gov-west-1a", false},
		{"Invalid Region Name", "us-east-coast-1", false},
		{"Valid Region with 3-digit number", "ap-southeast-123", true},
		{"Invalid Region with 4-digit number", "eu-west-1234", false},
		{"Invalid Region with uppercase letters", "US-WEST-2", false},
		{"Invalid Region with special characters", "eu-west-1!", false},
		{"Valid Region with 2-letter country code", "af-south-1", true},
		{"Invalid Region with 1-letter country code", "a-west-1", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validate.LikeAWSRegion(tc.input)
			if result != tc.expected {
				t.Errorf("expected '%v', but got '%v' for '%s'", tc.expected, result, tc.name)
			}
		})
	}
}

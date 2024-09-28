package fp_cm_utils

import (
	"testing"
)

func TestMaskEmailAddressInformation(t *testing.T) {
	tests := []struct {
		email         string
		maskBy        string
		revealCounter int
		expected      string
		expectErr     bool
	}{
		{"user@example.com", "*", 2, "us**@example.com", false},
		{"user@example.com", "#", 3, "use#@example.com", false},
		{"user@example.com", "*", 5, "user@example.com", false},
		{"short@domain.com", "!", 3, "sho!!@domain.com", false},
		{"short.jhon.hn@domain.com", "!", 5, "short!!!!!!!!@domain.com", false},
		{"a@b.com", "$", 1, "a@b.com", false},
		{"invalidemail.com", "*", 2, "", true},                 // Invalid email format
		{"@domain.com", "*", 2, "", true},                      // Invalid email format
		{"user@.com", "*", 2, "", true},                        // Invalid email format
		{"user@example.com", "", 2, "us**@example.com", false}, // maskBy is empty
		{"user@example.com", "**", 2, "", true},                // maskBy is not a single character
		{"user@example.com", "*", 10, "user@example.com", false},
	}

	for _, test := range tests {
		result, err := MaskEmailAddressInformation(test.email, test.maskBy, test.revealCounter)

		if (err != nil) != test.expectErr {
			t.Errorf("email: %s, maskBy: %s, revealCount: %d - expected error: %v, got: %v", test.email, test.maskBy, test.revealCounter, test.expectErr, err)
		}
		if result != test.expected {
			t.Errorf("email: %s, maskBy: %s, revealCount: %d - expected: %s, got: %s", test.email, test.maskBy, test.revealCounter, test.expected, result)
		}
	}
}

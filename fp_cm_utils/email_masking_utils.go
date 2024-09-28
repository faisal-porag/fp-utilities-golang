package fp_cm_utils

import (
	"errors"
	"strings"
)

// MaskEmailAddressInformation masks the email address, allowing dynamic replacement of characters.
// It keeps the specified number of characters in the local part visible.
// If revealCount exceeds the local part length, the original email is returned.
// Returns an error if the input email is not valid.
func MaskEmailAddressInformation(email string, maskBy string, revealCounter int) (string, error) {
	if !__isValidEmail(email) {
		return "", errors.New("invalid email address format")
	}
	if strings.TrimSpace(maskBy) == "" {
		maskBy = "*" // Set Default
	}
	if len(maskBy) != 1 {
		return "", errors.New("maskBy parameter must be a single character")
	}

	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]

	// Ensure revealCounter doesn't exceed the local part length
	if revealCounter > len(local) {
		revealCounter = len(local)
	}

	// Mask the local part, keeping the specified number of characters
	if len(local) > revealCounter {
		maskedLocal := local[:revealCounter] + strings.Repeat(maskBy, len(local)-revealCounter)
		return maskedLocal + "@" + domain, nil
	}

	return email, nil // Return original if local part is empty or shorter than revealCount
}

// __isValidEmail checks if the provided email has a valid format.
func __isValidEmail(email string) bool {
	// Simple validation for illustrative purposes.
	// Check if there's exactly one '@' and that both parts are well-formed.
	if strings.Count(email, "@") != 1 {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}

	// Check if the domain part contains at least one '.' and is not starting or ending with '.'.
	domain := parts[1]
	if !strings.Contains(domain, ".") || strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}

	return true
}

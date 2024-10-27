package fp_cm_utils

import (
	"fmt"
	"regexp"
	"strings"
)

func SimpleAddressValidate(input string) error {
	pattern := `^[a-zA-Z0-9\s#.&,'\-:/_()@:]+$`
	validCharRegex := regexp.MustCompile(pattern)

	// Collect invalid characters
	var invalidChars []string
	for _, char := range input {
		if !validCharRegex.MatchString(string(char)) {
			invalidChars = append(invalidChars, string(char))
		}
	}

	// If invalid characters found, create an error message
	if len(invalidChars) > 0 {
		return fmt.Errorf("Invalid address input: the following characters are not allowed: (%s)", strings.Join(invalidChars, ", "))
	}

	return nil
}


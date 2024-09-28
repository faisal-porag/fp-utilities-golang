package fp_cm_utils

import (
	"testing"
	"time"
)

func TestGetMultiLanguageFormatDate(t *testing.T) {
	date := time.Date(2023, time.January, 2, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		lang     string
		expected string
		hasError bool
	}{
		{"en", "Mon, Jan 2, 2023", false},
		{"bn", "সোম. ২ জানুয়ারী ২০০৬", false},
		{"france", "lun. 2 janv. 2023", false},
		{"german", "Mo. 2. Jan. 2023", false},
		{"china", "2006年1月2日周一", false},
		{"japan", "2006年1月2日月", false},
		{"hindi", "सोम. 2 जन. 2006", false},
		{"es", "", true}, // Unsupported language
	}

	for _, test := range tests {
		t.Run(test.lang, func(t *testing.T) {
			result, err := GetMultiLanguageFormatDate(date, test.lang)

			if test.hasError {
				if err == nil {
					t.Errorf("Expected an error for language %s, got none", test.lang)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error for language %s, got: %v", test.lang, err)
				}
				if result != test.expected {
					t.Errorf("For language %s, expected %v, got %v", test.lang, test.expected, result)
				}
			}
		})
	}
}

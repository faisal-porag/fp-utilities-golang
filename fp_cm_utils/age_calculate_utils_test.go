package fp_cm_utils

import (
	"testing"
)

func TestCalculateAge(t *testing.T) {
	calculator := AgeCalculator{}

	tests := []struct {
		dob         string
		expected    Age
		expectError bool
	}{
		{"1990-05-15", Age{34, 4, 14}, false},
		{"2000/01/01", Age{24, 8, 30}, false},
		{"invalid-date", Age{}, true},
	}

	for _, test := range tests {
		age, err := calculator.Calculate(test.dob)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected an error for DOB: %s, got none", test.dob)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect an error for DOB: %s, got: %v", test.dob, err)
			}
			if age != test.expected {
				t.Errorf("For DOB: %s, expected age: %+v, got: %+v", test.dob, test.expected, age)
			}
		}
	}
}

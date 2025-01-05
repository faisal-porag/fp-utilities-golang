package fp_cm_utils

import (
	"testing"
)

// TestGenerateNumberToWordFunction tests the GenerateNumberToWordFunction.
func TestGenerateNumberToWordFunction(t *testing.T) {
	numberToWord := GenerateNumberToWordFunction()

	tests := []struct {
		input    int
		expected string
	}{
		{0, "Zero"},
		{5, "Five"},
		{10, "Ten"},
		{15, "Fifteen"},
		{20, "Twenty"},
		{25, "Twenty Five"},
		{100, "One Hundred"},
		{105, "One Hundred Five"},
		{123, "One Hundred Twenty Three"},
		{1000, "One Thousand"},
		{1010, "One Thousand Ten"},
		{1100, "One Thousand One Hundred"},
		{1123, "One Thousand One Hundred Twenty Three"},
		{1000000, "One Million"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{1000000000, "One Billion"},
		{1000000001, "One Billion One"},
	}

	for _, test := range tests {
		result := numberToWord(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %q but got %q", test.input, test.expected, result)
		}
	}
}

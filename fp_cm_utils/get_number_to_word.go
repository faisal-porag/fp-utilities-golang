package fp_cm_utils

import "strings"

// **********************************
// Note:: How to use
//numberToWord := utils.GenerateNumberToWordFunction()
//numberToWordRes := numberToWord(25)
//fmt.Println(numberToWordRes)
//************************************

// GenerateNumberToWordFunction >>>> Function to generate the number to word function dynamically
func GenerateNumberToWordFunction() func(int) string {
	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	scales := []string{"", "Thousand", "Million", "Billion"}

	// Return a dynamically constructed function
	return func(num int) string {
		if num == 0 {
			return "Zero"
		}

		// Break the number into groups of three digits
		var parts []string
		scaleIndex := 0

		for num > 0 {
			if num%1000 != 0 {
				parts = append([]string{___convertThreeDigits(num%1000, ones, teens, tens)}, parts...)
				if scaleIndex > 0 {
					parts[0] += " " + scales[scaleIndex]
				}
			}
			num /= 1000
			scaleIndex++
		}

		return strings.Join(parts, " ")
	}
}

// Helper function to convert numbers up to 999
func ___convertThreeDigits(num int, ones, teens, tens []string) string {
	hundreds := num / 100
	remainder := num % 100
	var parts []string

	if hundreds > 0 {
		parts = append(parts, ones[hundreds]+" Hundred")
	}
	if remainder >= 10 && remainder < 20 {
		parts = append(parts, teens[remainder-10])
	} else {
		if remainder >= 20 {
			parts = append(parts, tens[remainder/10])
		}
		if remainder%10 > 0 {
			parts = append(parts, ones[remainder%10])
		}
	}

	return strings.Join(parts, " ")
}

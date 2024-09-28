package fp_cm_utils

import (
	"fmt"
	"time"
)

// Age represents the age in years, months, and days.
type Age struct {
	Years  int
	Months int
	Days   int
}

type AgeCalculator struct{}

// Calculate calculates the age based on the given date of birth.
func (a *AgeCalculator) Calculate(dob string) (Age, error) {
	birthDate, err := parseDOB(dob)
	if err != nil {
		return Age{}, err
	}

	currentDate := time.Now()
	years, months, days := __calculateAge(birthDate, currentDate)

	return Age{Years: years, Months: months, Days: days}, nil
}

// parseDOB attempts to parse a date of birth string using multiple layouts.
// It returns the parsed time.Time and any parsing error encountered.
func parseDOB(dob string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",      // ISO format
		"02-01-2006",      // DD-MM-YYYY format
		"01-02-2006",      // MM-DD-YYYY format
		"2006/01/02",      // ISO format with slashes
		"02/01/2006",      // DD-MM-YYYY format with slashes
		"01/02/2006",      // MM-DD-YYYY format with slashes
		"January 2, 2006", // Full month name
		"Jan 2, 2006",     // Abbreviated month name
		"2 Jan 2006",      // Day first with month
		"2 January 2006",  // Day first with full month
	}

	for _, layout := range layouts {
		birthDate, err := time.Parse(layout, dob)
		if err == nil {
			return birthDate, nil // Return the parsed date if successful
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s. please pass a valid date format", dob)
}

// __calculateAge calculates the difference between two dates.
func __calculateAge(birthDate, currentDate time.Time) (years, months, days int) {
	years = currentDate.Year() - birthDate.Year()
	months = int(currentDate.Month()) - int(birthDate.Month())
	days = currentDate.Day() - birthDate.Day()

	if days < 0 {
		months-- // Adjust for negative days
		// Calculate days in the previous month
		previousMonth := currentDate.AddDate(0, -1, 0)
		days += __daysInMonth(previousMonth.Year(), int(previousMonth.Month()))
	}

	if months < 0 {
		years-- // Adjust for negative months
		months += 12
	}

	return years, months, days
}

// __daysInMonth returns the number of days in a given month and year.
func __daysInMonth(year, month int) int {
	if month == 2 {
		if __isLeapYear(year) {
			return 29
		}
		return 28
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}

// __isLeapYear checks if a year is a leap year.
func __isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

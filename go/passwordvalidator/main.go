package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ValidatePassword(input string) (bool, error) {
	errorMessages := make([]string, 0, 2)

	if err := validateLength(input); len(err) > 0 {
		errorMessages = append(errorMessages, err)
	}

	if err := validateNumberOfDigits(input); len(err) > 0 {
		errorMessages = append(errorMessages, err)
	}

	if len(errorMessages) > 0 {
		return false, buildErrorOutput(errorMessages)
	}
	return true, nil
}

func validateLength(input string) string {
	if len(input) < 8 {
		return "password must be at least 8 characters"
	}
	return ""
}

func validateNumberOfDigits(input string) string {
	numberOfDigits := 0
	for _, element := range input {
		if unicode.IsDigit(element) {
			numberOfDigits += 1
		}
	}

	if numberOfDigits < 2 {
		return "the password must contain at least 2 numbers"
	}

	return ""
}

func buildErrorOutput(errorMessages []string) error {
	return fmt.Errorf("%s", strings.Join(errorMessages, "\n"))
}

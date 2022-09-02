package main

import (
	"fmt"
	"strings"
	"unicode"
)

type validationStep func(string) string

func ValidatePassword(input string) (bool, error) {
	errorMessages := make([]string, 0, 2)

	errorMessages = runValidationStep(validateLength, input, errorMessages)
	errorMessages = runValidationStep(validateNumberOfDigits, input, errorMessages)
	errorMessages = runValidationStep(validateOneCapitalLetter, input, errorMessages)
	errorMessages = runValidationStep(validateOneSpecialCharacter, input, errorMessages)

	if len(errorMessages) > 0 {
		return false, buildErrorOutput(errorMessages)
	}
	return true, nil
}

func runValidationStep(step validationStep, argument string, errList []string) []string {
	if err := step(argument); len(err) > 0 {
		return append(errList, err)
	}
	return errList
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

func validateOneCapitalLetter(input string) string {
	for _, element := range input {
		if unicode.IsUpper(element) {
			return ""
		}
	}
	return "password must contain at least one capital letter"
}

func validateOneSpecialCharacter(input string) string {
	for _, element := range input {
		if !unicode.IsDigit(element) && !unicode.IsLetter(element) {
			return ""
		}
	}
	return "password must contain at least one special character"
}

func buildErrorOutput(errorMessages []string) error {
	return fmt.Errorf("%s", strings.Join(errorMessages, "\n"))
}

package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type StringCalculatorState int64

const STANDARDIZED_SEPARATOR string = ","
const NEWLINE string = "\n"

const (
	EXPECTING_DIGIT              StringCalculatorState = 0
	EXPECTING_SEPARATOR          StringCalculatorState = 1
	EXPECTING_DIGIT_OR_SEPARATOR StringCalculatorState = 2
)

type StringCalculator struct {
	Input                  string
	standardizedInput      string
	extractedDigits        []int
	currentExtractionState StringCalculatorState
}

func (sc *StringCalculator) Add(numbers string) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	} else {
		error_msg := ""
		sc.Input = numbers
		sc.currentExtractionState = EXPECTING_DIGIT
		sc.standardizeInput()
		error_msg += sc.validateEndOfInput()
		sc.extractDigits()

		if len(error_msg) > 0 {
			return -1, errors.New(error_msg)
		}
		return sc.sumDigits(), nil
	}
}

func (sc *StringCalculator) standardizeInput() {
	if strings.HasPrefix(sc.Input, "//") {
		sep := sc.extractCustomSeparator()
		indexOfStringStart := strings.Index(sc.Input, "\n")
		sc.standardizedInput = strings.Replace(sc.Input[indexOfStringStart:], sep, STANDARDIZED_SEPARATOR, -1)
	} else {
		sc.standardizedInput = strings.Replace(sc.Input, "\n", STANDARDIZED_SEPARATOR, -1)
	}
}

func (sc StringCalculator) extractCustomSeparator() string {
	sep := ""
	for _, element := range sc.Input[2:] {
		element_as_string := string(element)
		if element_as_string == NEWLINE {
			break
		}
		sep += element_as_string
	}
	return sep
}

func (sc StringCalculator) replaceSeparatorOnInput(sep string, indexOfStringStart int) {

}

func (sc StringCalculator) validateEndOfInput() string {
	if string(sc.Input[len(sc.Input)-1]) == STANDARDIZED_SEPARATOR {
		return "Cannot have a separator at the end of the input"
	}
	return ""
}

func (sc *StringCalculator) extractDigits() {
	current_digit := ""
	for _, element := range sc.standardizedInput {
		if unicode.IsDigit(element) && (sc.currentExtractionState == EXPECTING_DIGIT || sc.currentExtractionState == EXPECTING_DIGIT_OR_SEPARATOR) {
			current_digit += string(element)
			sc.currentExtractionState = EXPECTING_DIGIT_OR_SEPARATOR
		} else if string(element) == STANDARDIZED_SEPARATOR {
			if res, err := strconv.Atoi(current_digit); err == nil {
				current_digit = ""
				sc.extractedDigits = append(sc.extractedDigits, res)
				sc.currentExtractionState = EXPECTING_DIGIT
			}
		}
	}

	if len(current_digit) > 0 {
		if res, err := strconv.Atoi(current_digit); err == nil {
			sc.extractedDigits = append(sc.extractedDigits, res)
		}
	}
}

func (sc StringCalculator) sumDigits() int {
	result := 0
	for _, v := range sc.extractedDigits {
		result += v
	}
	return result
}

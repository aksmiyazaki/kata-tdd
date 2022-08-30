package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type StringCalculatorState int64

const NEWLINE string = "\n"

const (
	EXPECTING_DIGIT              StringCalculatorState = 0
	EXPECTING_SEPARATOR          StringCalculatorState = 1
	EXPECTING_DIGIT_OR_SEPARATOR StringCalculatorState = 2
)

type StringCalculator struct {
	Input                    string
	standardizedInput        string
	extractedDigits          []int
	currentStandardSeparator string
	originalSeparators       []string
}

func (sc *StringCalculator) Add(numbers string) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	} else {
		error_msg := ""
		sc.Input = numbers
		if err := sc.standardizeInput(); err != nil {
			return -1, err
		}

		error_msg += sc.validateEndOfInput()
		if err := sc.extractDigits(); err != nil {
			return -1, err
		}

		if len(error_msg) > 0 {
			return -1, errors.New(error_msg)
		}
		return sc.sumDigits(), nil
	}
}

func (sc *StringCalculator) standardizeInput() error {
	var err error
	err = nil

	if sc.isCustomSeparatedInput() {
		err = sc.standardizeCustomSeparatedInput()
	} else {
		sc.standardizeDefaultSeparatedInput()
	}

	return err
}

func (sc StringCalculator) isCustomSeparatedInput() bool {
	return strings.HasPrefix(sc.Input, "//")
}

func (sc *StringCalculator) standardizeCustomSeparatedInput() error {
	sep := sc.extractCustomSeparator()
	indexOfStringStart := strings.Index(sc.Input, "\n") + 1
	if standardSeparator, err := sc.checkStringForStandardSeparators(sc.Input[indexOfStringStart:]); err != nil {
		return err
	} else {
		sc.originalSeparators = append(sc.originalSeparators, sep)
		sc.currentStandardSeparator = standardSeparator
		sc.standardizedInput = strings.Replace(sc.Input[indexOfStringStart:], sep, sc.currentStandardSeparator, -1)
	}
	return nil
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

func (sc StringCalculator) checkStringForStandardSeparators(customizedInput string) (string, error) {
	options := sc.getStandardOperatorOptions()
	for _, element := range options {
		if !strings.Contains(customizedInput, element) {
			return element, nil
		}
	}
	return "", errors.New("cannot use a standardized operator")
}

func (sc StringCalculator) getStandardOperatorOptions() []string {
	return []string{",", "\n"}
}

func (sc *StringCalculator) standardizeDefaultSeparatedInput() {
	sc.originalSeparators = sc.getStandardOperatorOptions()
	sc.currentStandardSeparator = sc.getStandardOperatorOptions()[0]
	sc.standardizedInput = strings.Replace(sc.Input, "\n", sc.currentStandardSeparator, -1)
}

func (sc StringCalculator) validateEndOfInput() string {
	if string(sc.Input[len(sc.Input)-1]) == sc.currentStandardSeparator {
		return "Cannot have a separator at the end of the input"
	}
	return ""
}

func (sc *StringCalculator) extractDigits() error {
	current_digit := ""
	currentExtractionState := EXPECTING_DIGIT

	for idx, element := range sc.standardizedInput {
		elementAsString := string(element)
		if unicode.IsDigit(element) && (currentExtractionState == EXPECTING_DIGIT || currentExtractionState == EXPECTING_DIGIT_OR_SEPARATOR) {
			current_digit += elementAsString
			currentExtractionState = EXPECTING_DIGIT_OR_SEPARATOR
		} else if elementAsString == sc.currentStandardSeparator {
			if res, err := strconv.Atoi(current_digit); err == nil {
				current_digit = ""
				sc.extractedDigits = append(sc.extractedDigits, res)
				currentExtractionState = EXPECTING_DIGIT
			}
		} else {
			expectedSeparators := strings.Join(sc.originalSeparators, "','")
			return fmt.Errorf("'%s' expected but '%s' found at position %d", expectedSeparators, elementAsString, idx)
		}
	}

	if len(current_digit) > 0 {
		if res, err := strconv.Atoi(current_digit); err == nil {
			sc.extractedDigits = append(sc.extractedDigits, res)
		}
	}
	return nil
}

func (sc StringCalculator) sumDigits() int {
	result := 0
	for _, v := range sc.extractedDigits {
		result += v
	}
	return result
}

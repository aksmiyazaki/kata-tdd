package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		description    string
		input          string
		expectedOutput int
		expectedError  error
	}

	for _, testCase := range []testCase{
		{
			description:    "ShouldReturn0WhenInputIsEmpty",
			input:          "",
			expectedOutput: 0,
			expectedError:  nil,
		},
		{
			description:    "ShouldReturnInputWhenSingleDigit",
			input:          "1",
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			description:    "ShouldReturnSumWhenMultipleDigitsAreInput",
			input:          "1,2",
			expectedOutput: 3,
			expectedError:  nil,
		},
		{
			description:    "ShouldReturnSumWhenUsingNewlineAsSeparator",
			input:          "1,2\n3",
			expectedOutput: 6,
			expectedError:  nil,
		},
		{
			description:    "ShouldRaiseErrorWhenTheresSeparatorAtTheEnd",
			input:          "1,2,",
			expectedOutput: -1,
			expectedError:  errors.New("Cannot have a separator at the end of the input"),
		},
		{
			description:    "ShouldSumElementsWhenCustomDelimited",
			input:          "//;\n1;3",
			expectedOutput: 4,
			expectedError:  nil,
		},
		{
			description:    "ShouldSumElementsWhenCustomDelimitedWithMultipleDigitSeparator",
			input:          "//sep\n2sep3sep4",
			expectedOutput: 9,
			expectedError:  nil,
		},
		{
			description:    "ShouldRaiseErrorWhenUsingInvalidSeparator",
			input:          "//|\n1|2,3",
			expectedOutput: -1,
			expectedError:  errors.New("'|' expected but ',' found at position 3"),
		},
		{
			description:    "ShouldRaiseErrorWhenTheresNegatives",
			input:          "1,-2",
			expectedOutput: -1,
			expectedError:  errors.New("negative number(s) not allowed: -2"),
		},
		{
			description:    "ShouldRaiseErrorWithAllNegativesWhenTheresMultipleNegatives",
			input:          "//|\n1|-2|-3",
			expectedOutput: -1,
			expectedError:  errors.New("negative number(s) not allowed: -2,-3"),
		},
		{
			description:    "ShouldRaiseMultipleErrorsWhenTheresMultipleErrors",
			input:          "//|\n1|2,-3",
			expectedOutput: -1,
			expectedError:  errors.New("'|' expected but ',' found at position 3\nnegative number(s) not allowed: -3"),
		},
		{
			description:    "ShouldIgnoreNumberWhenItIsBiggerThan1000",
			input:          "2,1000,1001",
			expectedOutput: 1002,
			expectedError:  nil,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			s := StringCalculator{}
			output, err := s.Add(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
			assert.Equal(t, testCase.expectedError, err)
		})
	}

}

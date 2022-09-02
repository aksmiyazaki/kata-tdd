package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	type testCase struct {
		description    string
		input          string
		expectedError  error
		expectedOutput bool
	}

	for _, testCase := range []testCase{
		{
			description:    "ShouldFailWhenPasswordIsSmallerThan8Chars",
			input:          "A23456;",
			expectedError:  errors.New("password must be at least 8 characters"),
			expectedOutput: false,
		},
		{
			description:    "ShouldFailWhenThereIsLessThan2Numbers",
			input:          "1Abcdefg'",
			expectedError:  errors.New("the password must contain at least 2 numbers"),
			expectedOutput: false,
		},
		{
			description:    "ShouldReturnMultipleErrorsInASingleValidation",
			input:          "Somepa+",
			expectedError:  errors.New("password must be at least 8 characters\nthe password must contain at least 2 numbers"),
			expectedOutput: false,
		},
		{
			description:    "ShouldFailWhenThereIsntACapitalLetter",
			input:          "12somepas[",
			expectedError:  errors.New("password must contain at least one capital letter"),
			expectedOutput: false,
		},
		{
			description:    "ShouldValidateWhenPasswordMatchesAllConstraints",
			input:          "12somepaS[",
			expectedError:  nil,
			expectedOutput: true,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			output, errOutput := ValidatePassword(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
			assert.Equal(t, testCase.expectedError, errOutput)
		})
	}

}

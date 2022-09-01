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
			input:          "1234567",
			expectedError:  errors.New("password must be at least 8 characters"),
			expectedOutput: false,
		},
		{
			description:    "ShouldFailWhenThereIsLessThan2Numbers",
			input:          "1abcdefgh",
			expectedError:  errors.New("the password must contain at least 2 numbers"),
			expectedOutput: false,
		},
		{
			description:    "ShouldReturnMultipleErrorsInASingleValidation",
			input:          "somepas",
			expectedError:  errors.New("password must be at least 8 characters\nthe password must contain at least 2 numbers"),
			expectedOutput: false,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			output, errOutput := ValidatePassword(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
			assert.Equal(t, testCase.expectedError, errOutput)
		})
	}

}

package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosScan(t *testing.T) {
	type testCase struct {
		description    string
		input          string
		expectedOutput string
		expectedError  error
	}

	for _, testCase := range []testCase{
		{
			description:    "ShouldDisplayPriceWhenProductExists",
			input:          "12345",
			expectedOutput: "$7.25",
			expectedError:  nil,
		},
		{
			description:    "ShouldDisplayErrorWhenBarCodeIsNotFound",
			input:          "999999",
			expectedOutput: "",
			expectedError:  errors.New("barcode not found"),
		},
		{
			description:    "ShouldDisplayErrorWhenBarCodeIsEmpty",
			input:          "",
			expectedOutput: "",
			expectedError:  errors.New("empty barcode"),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			testingPOE := NewPOS()
			output, err := testingPOE.Scan(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestPosTotal(t *testing.T) {
	type testCase struct {
		description    string
		input          []string
		expectedOutput string
	}

	for _, testCase := range []testCase{
		{
			description:    "ShouldSumWhenProductsExists",
			input:          []string{"12345", "489", "555"},
			expectedOutput: "$27.25",
		},
		{
			description:    "ShouldPartiallySumWhenSomeProductsExists",
			input:          []string{"12345", "999", "555"},
			expectedOutput: "$17.25",
		},
		{
			description:    "ShouldReturnZeroWhenNoProductsExists",
			input:          []string{"888", "999"},
			expectedOutput: "$0.00",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			testingPOE := NewPOS()
			output := testingPOE.Total(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type testCase struct {
		description    string
		input          string
		expectedOutput []string
	}

	for _, testCase := range []testCase{
		{
			description:    "ShouldFailWhenInputIsFewerThan2Characters",
			input:          "V",
			expectedOutput: []string{},
		},
		{
			description:    "ShouldReturnAllMatchingStringStartsWhenInputIsBiggerThan1Character",
			input:          "Va",
			expectedOutput: []string{"Valencia", "Vancouver"},
		},
		{
			description:    "ShouldIgnoreCaseWhenInputIsDifferentThanDatabase",
			input:          "VA",
			expectedOutput: []string{"Valencia", "Vancouver"},
		},
		{
			description:    "ShouldMatchWhenMatchIsAnyPartOfString",
			input:          "ape",
			expectedOutput: []string{"Budapest"},
		},
		{
			description: "ShouldReturnAllDatabaseWhenQueryIsWildcard",
			input:       "*",
			expectedOutput: []string{"Paris", "Budapest", "Skopje",
				"Rotterdam", "Valencia", "Vancouver",
				"Amsterdam", "Vienna", "Sydney",
				"New York City", "London", "Bangkok",
				"Hong Kong", "Dubai", "Rome", "Istanbul"},
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			output := Search(testCase.input)
			assert.Equal(t, testCase.expectedOutput, output)
		})
	}

}

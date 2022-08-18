package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturn0WhenInputIsEmpty(t *testing.T) {
	res := Add("")
	assert.Equal(t, 0, res)
}

func TestShouldReturnInputWhenSingleDigit(t *testing.T) {
	res := Add("1")
	assert.Equal(t, 1, res)
}

func TestShouldReturnSumWhenMultipleDigits(t *testing.T) {

}

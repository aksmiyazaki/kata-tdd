package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNumberAsString(t *testing.T) {
	res := fizzBuzz(1)
	assert.Equal(t, "1", res)
}

func TestShouldReturnFizzWhenInputIsMultipleOfThree(t *testing.T) {
	res := fizzBuzz(3)
	assert.Equal(t, "Fizz", res)
}

func TestShouldReturnBuzzWhenInputIsMultipleOfFive(t *testing.T) {
	res := fizzBuzz(5)
	assert.Equal(t, "Buzz", res)
}

func TestShouldReturnFizzBuzzWhenInputIsMultipleOfThreeAndFive(t *testing.T) {
	res := fizzBuzz(15)
	assert.Equal(t, "FizzBuzz", res)
}

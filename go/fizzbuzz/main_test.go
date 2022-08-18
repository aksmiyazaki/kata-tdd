package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNumberAsStringWhenItIsNotMultipleOf3Or5(t *testing.T) {
	res := fizzBuzz(1)
	assert.Equal(t, "1", res)
}

func TestShouldReturnFizzWhenInputIsMultipleOf3(t *testing.T) {
	res := fizzBuzz(3)
	assert.Equal(t, "Fizz", res)
}

func TestShouldReturnBuzzWhenInputIsMultipleOf5(t *testing.T) {
	res := fizzBuzz(5)
	assert.Equal(t, "Buzz", res)
}

func TestShouldReturnFizzBuzzWhenInputIsMultipleOf3And5(t *testing.T) {
	res := fizzBuzz(15)
	assert.Equal(t, "FizzBuzz", res)
}

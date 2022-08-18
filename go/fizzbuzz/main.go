package main

import "strconv"

func fizzBuzz(input int) string {
	output := ""

	if input%3 == 0 {
		output += "Fizz"
	}
	if input%5 == 0 {
		output += "Buzz"
	}

	if len(output) == 0 {
		return strconv.Itoa(input)
	} else {
		return output
	}
}

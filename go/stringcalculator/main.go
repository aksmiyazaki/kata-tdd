package main

import (
	"strconv"
)

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else {
		res, _ := strconv.Atoi(numbers)
		return res
	}
}

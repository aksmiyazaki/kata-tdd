package main

import "strings"

func Search(input string) []string {
	if input == "*" {
		return getCityDatabase()
	}

	if len(input) < 2 {
		return []string{}
	}

	return matchInputWithCityDatabase(input)
}

func getCityDatabase() []string {
	return []string{"Paris", "Budapest", "Skopje",
		"Rotterdam", "Valencia", "Vancouver",
		"Amsterdam", "Vienna", "Sydney",
		"New York City", "London", "Bangkok",
		"Hong Kong", "Dubai", "Rome", "Istanbul"}
}

func matchInputWithCityDatabase(input string) []string {
	matchingValues := []string{}
	caseInsensitiveInput := strings.ToUpper(input)
	for _, el := range getCityDatabase() {
		caseInsensitiveEl := strings.ToUpper(el)

		if strings.Contains(caseInsensitiveEl, caseInsensitiveInput) {
			matchingValues = append(matchingValues, el)
		}
	}
	return matchingValues
}

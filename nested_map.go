package main

import (
	"unicode/utf8"
)

func getNameCounts(names []string) map[rune]map[string]int {
	// takes slice of strings, names -> returns nested map
	// parent map's keys are unique first characters

	// billy
	// map["b"]map["billy"]2

	nestedMap := make(map[rune]map[string]int)

	// every name in the slice, ommit index
	for _, name := range names {
		// ignore names of 0 characters
		if len(name) == 0 {
			continue
		}

		// first rune is key of parent dict
		// rune := rune(name[0]) doesn't handle emojis
		rune, _ := utf8.DecodeRuneInString(name)

		// check if the first char's map exists, make it otherwise
		if _, ok := nestedMap[rune]; !ok {
			nestedMap[rune] = make(map[string]int)
		}
		// increment count
		nestedMap[rune][name] += 1
	}
	return nestedMap
}

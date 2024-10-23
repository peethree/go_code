package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	// go over all the strings in message

	// each distinct word count + 1

	// add strings to dict, then return key count

	distinctWords := make(map[string]int)

	// for each msg in the messages list
	for _, message := range messages {
		// split message into individual strings, and make the words lower case
		individualWords := strings.Fields(strings.ToLower(message))

		// for every word in the split message
		for _, word := range individualWords {
			// add that word to dict
			distinctWords[word] = 0
		}
	}
	// return the amount of keys in the map, which should be the same as distinct # of words
	return len(distinctWords)

}

package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	// slice to hold the slices of bytes
	sliceofSlices := [][]byte{}

	// loop over every item
	for i := 0; i < len(items); i++ {
		// marshal the json representation into bytes
		data, err := json.Marshal(items[i])
		if err != nil {
			return [][]byte{}, err
		}
		// append to slice
		sliceofSlices = append(sliceofSlices, data)
	}

	return sliceofSlices, nil
}

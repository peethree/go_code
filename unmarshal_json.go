package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Change the return signature to return []Item instead of []byte.
func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return []Item{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Get the data from the response body using io.ReadAll, creating a slice of bytes []byte.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []Item{}, err
	}

	// Create a nil slice of items []Item.
	var items []Item
	// Use json.Unmarshal on the data to get the JSON data.
	if err = json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}
	// Return the items.
	return items, nil
}

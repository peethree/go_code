package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	res, err := http.Get(url)
	// if network error occurs
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}

	defer res.Body.Close()

	// If the response does not have status code 200 return the status code, and the following error, where %s is the .Status field of the response.
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}
	// If no errors occurred simply return the .StatusCode and nil.
	return res.StatusCode, nil
}

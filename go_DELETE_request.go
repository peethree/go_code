package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// new request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// modify req headers
	req.Header.Set("X-API-Key", apiKey)

	// make the request with do method
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// resource cleanup
	defer res.Body.Close()

	// check status code
	if res.StatusCode > 299 {
		return errors.New("request to delete user unsuccessful")
	}

	return nil
}

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []User{}, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("Character name: %s, Class: %s, Level: %d, User: %s\n", user.CharacterName, user.Class, user.Level, user.User.Name)
	}
}

type User struct {
	Id            string `json:"id"`
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	Level         int    `json:"level"`
	User          struct {
		Name string `json:"name"`
	} `json:"user"`
}

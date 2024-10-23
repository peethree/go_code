package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// encode data as json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// new req, use a bytes.NewBuffer to create a io.Reader from the JSON data.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// modify headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// make new client
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	defer res.Body.Close()

	// decode n return json body / User
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil

}

type User struct {
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	Level         int    `json:"level"`
	PvpEnabled    bool   `json:"pvpEnabled"`
	User          struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

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

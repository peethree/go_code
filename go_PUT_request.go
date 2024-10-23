package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	// encode as json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// new PUT request
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

	// modify req headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// make the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	// decode and reutnr response's JSON body (User)
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	// modify header
	req.Header.Set("X-API-Key", apiKey)

	// make request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	defer res.Body.Close()

	// decode
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
	ID            string `json:"id"`
	Level         int    `json:"level"`
	PvpEnabled    bool   `json:"pvpEnabled"`
	User          struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUser(user User) {
	fmt.Printf("User uuid: %s, Character Name: %s, Class: %s, Level: %d, PVP Status: %t, User name: %s\n",
		user.ID, user.CharacterName, user.Class, user.Level, user.PvpEnabled, user.User.Name)
}

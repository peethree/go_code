package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(url string) ([]User, error) {

	var users []User

	// Call http.Get using its url parameter.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Decode the JSON data from the response and return it.
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&users); err != nil {
		return users, err
	}

	return users, nil
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

func logUsers(characters []User) {
	for _, char := range characters {
		fmt.Printf("Character name: %s, Class: %s, Level: %d, User: %s\n", char.CharacterName, char.Class, char.Level, char.User.Name)
	}
}

// // simple get request
// resp, err := http.Get("https://jsonplaceholder.typicode.com/users")

// // else: if there is a need to customize headers/ cookies/ timeouts
// client := &http.Client{
// 	Timeout: time.Second * 10,
// }

// req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
// if err != nil {
// 	log.Fatal(err)
// }

// resp, err := client.Do(req)

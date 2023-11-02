package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type UsernameAndPassword struct {
	Username string
	Password string
}
type Payload struct {
	Site        string
	Credentials *UsernameAndPassword
}

func WriteJSON(payload *Payload) {
	toWrite := &Payload{
		Site: payload.Site,
		Credentials: &UsernameAndPassword{
			Username: payload.Credentials.Username,
			Password: payload.Credentials.Password,
		},
	}

	file, err := os.Create("data.json")

	if err != nil {
		fmt.Println("Error creating the file:", err)

		return
	}

	defer file.Close()

	// Create an file encoder.
	encoder := json.NewEncoder(file)

	// Encode the data and write it to the file.
	err = encoder.Encode(toWrite)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)

		return
	}
}

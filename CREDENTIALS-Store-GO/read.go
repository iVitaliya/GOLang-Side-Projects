package main

import (
	"encoding/json"
	"errors"
	"os"
)

// Create a map to store the JSON data
var JsonData map[string]*UsernameAndPassword = make(map[string]*UsernameAndPassword)

type readData struct {
	Key   string
	Value *UsernameAndPassword
}

func ReadJSON(site string) (*readData, error) {
	// Open the JSON file for reading
	file, err := os.Open("data.json")
	if err != nil {
		return &readData{
			Key: site,
			Value: &UsernameAndPassword{
				Username: "",
				Password: "",
			},
		}, errors.New("Error opening the file: " + err.Error())
	}

	defer file.Close()

	// Create a decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the map.
	err = decoder.Decode(&JsonData)
	if err != nil {
		return &readData{
			Key: site,
			Value: &UsernameAndPassword{
				Username: "",
				Password: "",
			},
		}, errors.New("Error decoding JSON:" + err.Error())
	}

	// Now you can access the data in the map
	for key, val := range JsonData {
		if key == site {
			return &readData{
				Key: site,
				Value: &UsernameAndPassword{
					Username: val.Username,
					Password: val.Password,
				},
			}, nil
		}
	}

	return &readData{
		Key: site,
		Value: &UsernameAndPassword{
			Username: "",
			Password: "",
		},
	}, errors.New("The data that should've read " + site + " couldn't be found so we've returned nil/null")
}

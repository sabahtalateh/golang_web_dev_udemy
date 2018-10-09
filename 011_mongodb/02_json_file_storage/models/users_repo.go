package models

import (
	"encoding/json"
	"os"
)

// StoreUsers Stores users to a file
func StoreUsers(m map[string]User) {
	f, err := os.Create("storage/db.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(m)
	if err != nil {
		panic(err)
	}
}

// LoadUsers Loads users from file
func LoadUsers() map[string]User {
	f, err := os.Open("storage/db.json")
	if err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			f, err = os.Create("storage/db.json")
			if err != nil {
				panic(err)
			}
			m := make(map[string]User)
			err := json.NewEncoder(f).Encode(m)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	f, err = os.Open("storage/db.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := make(map[string]User)
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		panic(err)
	}
	return m
}

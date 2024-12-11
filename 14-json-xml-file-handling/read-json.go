package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// read the json, written at data.json file
// use json.Unmarshal to convert the byte data to a struct
// os.ReadFile, Scanner, (os.OpenFile -> f.Read)
// struct fields must be exported, so json package can work on it

// convert json to struct
// turn inline off to create different types for nested json
// https://mholt.github.io/json-to-go/
type user struct {
	FirstName    string          `json:"first_name"`    // json is a field level tag, used by the json package
	PasswordHash string          `json:"password_hash"` // setting name of the field in the json output
	Perms        map[string]bool `json:"perms"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Use a scanner to read the file
	scanner := bufio.NewScanner(file)
	var dataBytes []byte

	for scanner.Scan() {
		dataBytes = append(dataBytes, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse JSON data
	var users []user
	err = json.Unmarshal(dataBytes, &users)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Print the parsed data
	fmt.Println("Parsed data:")
	for _, user := range users {
		fmt.Printf("user: %+v\n", user)
	}
}

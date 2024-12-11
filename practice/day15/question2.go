// read the json, written at data.json file
// use json.Unmarshal to convert the byte data to a struct
// os.ReadFile, Scanner, (os.OpenFile -> f.Read)

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Permissions struct {
	Admin bool `json:"admin"`
}

type User struct {
	PasswordHash string      `json:"password_hash"`
	Perms        Permissions `json:"perms"`
}

func readWithScanner(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return []byte(sb.String()), nil
}

func main() {
	filePath := "data.json"

	data, err := readWithScanner(filePath)
	if err != nil {
		fmt.Println("Error reading file with bufio.Scanner:", err)
		return
	}
	fmt.Println("bufio.Scanner data:", string(data))

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, user := range users {
		fmt.Printf("Unmarshaled User: %+v\n", user)
	}
}

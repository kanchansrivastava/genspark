package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func main() {
	err := Handlers()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Handlers() error {
	f, err := OpenFile("file.txt")
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			log.Println(err)
			return fmt.Errorf("please create the file first then call the api")
		}
		return err
	}
	f.Close()
	return nil

}

func OpenFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
	}
	return f, nil
}

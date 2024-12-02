package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := OpenFile("test.txt")
	if err != nil {
		log.Println(err)
	}
	info, _ := f.Stat()
	fmt.Println(info.Name())
}

func OpenFile(name string) (*os.File, error) {
	f, err := os.Open(name)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(name)
			if err != nil {
				return nil, err
			}
			return f, nil
		}
		return nil, err
	}

	return f, nil
}

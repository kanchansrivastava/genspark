package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := openFile("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	info, _ := f.Stat()
	fmt.Println(info.Name())

}

func openFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("file %s does not exist\n", name)
			f, err := os.Create(name)
			if err != nil {
				return nil, fmt.Errorf("cannot create file %s: %w", name, err)
			}

			return f, nil
		}
		return nil, fmt.Errorf("file opening failed %s: %w", name, err)
	}

	return f, nil
}

// q5. {
// 	"name": "Alice",
// 	"age": 25
//   }

//   C:\Users\Alice\Documents\example.txt

//   Store above data in string

package main

import "fmt"

func main() {
	// Storing JSON data and file path as a string
	data := `{
		"name": "Alice",
		"age": 25
	}`
	filePath := `C:\Users\Alice\Documents\example.txt`

	// Print the data
	fmt.Println("JSON Data:", data)
	fmt.Println("File Path:", filePath)
}

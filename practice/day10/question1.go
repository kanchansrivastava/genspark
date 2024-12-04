
/* q1.	create a program that manages a collection of books and number of books available
    allow users to search for books by title.
	The program should handle errors gracefully if a book is not found or if there are any issues accessing the collection.


	Use a map to store book Name and their counter.
	Functionality:
	Implement
		- AddBook(title string,counter int) error
			-to add a new book to the collection.
	FetchBookCounter(name) (int, error)
		-to retrieve a book by its name.
	Error Handling:
	Use a struct to handle error
	User errors.As in main to check if struct is present inside the chain or not
*/
package main

import (
	"errors"
	"fmt"
)

var books = make(map[string]int)

type BookExistError struct {
	Func  string
	Input string
	Err   error
}

type BookNotFoundError struct {
	Func  string
	Input string
	Err   error
}

type CollectionAccessError struct {
	Func  string
	Input string
	Err   error
}

func (q *BookNotFoundError) Error() string {
	return fmt.Sprintf("main.%s: book '%s' not found: %v", q.Func, q.Input, q.Err)
}

func (q *BookExistError) Error() string {
	return fmt.Sprintf("main.%s: book '%s' already exists: %v", q.Func, q.Input, q.Err)
}

func (q *CollectionAccessError) Error() string {
	return fmt.Sprintf("main.%s: failed to access collection for '%s': %v", q.Func, q.Input, q.Err)
}

var ErrBookNotFound = errors.New("book not found")

var ErrCollectionAccess = errors.New("issue in accessing collection")


func AddBook(name string, counter int) error {
	if name == "" {
		return &CollectionAccessError{
			Func:  "AddBook",
			Input: name,
			Err:   errors.New("book title cannot be empty"),
		}
	}
	if counter <= 0 {
		return &CollectionAccessError{
			Func:  "AddBook",
			Input: name,
			Err:   errors.New("counter must be greater than 0"),
		}
	}
	// Check if book already exists
	if _, exists := books[name]; exists {
		return &BookExistError{
			Func:  "AddBook",
			Input: name,
			Err:   errors.New("book already exists"),
		}
	}
	books[name] = counter
	return nil
}

func FetchBookCounter(name string) (int, error) {
	counter, exists := books[name]
	if !exists {
		return 0, &BookNotFoundError{
			Func:  "FetchBookCounter",
			Input: name,
			Err:   ErrBookNotFound,
		}
	}
	return counter, nil
}

func main() {
	
	err := AddBook("Inner Child", 2) // Add a book to the collection
	if err != nil {
		var cae *CollectionAccessError
		if errors.As(err, &cae) {
			fmt.Println("Collection Access Error:", cae)
			return
		}

		var bne *BookNotFoundError
		if errors.As(err, &bne) {
			fmt.Println("Book Not Found Error:", bne)
			return
		}

		var be *BookExistError
		if errors.As(err, &be) {
			fmt.Println("Book Exists Error:", be)
			return
		}

		// Handle other types of errors
		fmt.Println("Unknown Error:", err)
		return
	}

	// Try fetching the book count
	counter, err := FetchBookCounter("Inner Childdd")
	if err != nil {
		var bne *BookNotFoundError
		if errors.As(err, &bne) {
			fmt.Println("Book Not Found Error:", bne)
			return
		}
		fmt.Println("Error fetching book:", err)
		return
	}
	fmt.Printf("Book 'Inner Child' has %d copies available\n", counter)
}

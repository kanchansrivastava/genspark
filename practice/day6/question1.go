/*
q1. Create a struct (Author)
    Two Field:- Name, Books[slice]

    Create two methods, one appends new books to the book slice , other prints the struct

    Create a function that accepts the struct and append values to the book slice

    Create a function that would accept the Books field, not the struct and append some more books
*/

package main

import "fmt"


type Author struct {
	Name  string
	Books []string
}

func (a *Author) AddBook(book string) {
	a.Books = append(a.Books, book) // directly updating a block x
}

func (a *Author) PrintAuthor() {
	fmt.Printf("Author: %s\nBooks: %v\n", a.Name, a.Books)
}

func AppendToAuthorBooks(author *Author, book string) {
	author.Books = append(author.Books, book)
}

func AppendToBooks(books *[]string, book string) {
	*books = append(*books, book)
}


func main() {
	author := Author{
		Name:  "John",
		Books: []string{"Book 1", "Book 2"},
	}

	author.AddBook("Book 3")
	author.PrintAuthor()

	AppendToAuthorBooks(&author, "Book 4")
	author.PrintAuthor()

	AppendToBooks(&author.Books, "Book 5")
	author.PrintAuthor()
}

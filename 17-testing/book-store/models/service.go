package models

import "context"

//go generate

// We write go generate to run the below line
// run go generate command from the current directory

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels
type Service interface {
	InsertBook(ctx context.Context, newBook NewBook) (Book, error)
	Update(ctx context.Context, id int, updateBook UpdateBook) (Book, error)
}

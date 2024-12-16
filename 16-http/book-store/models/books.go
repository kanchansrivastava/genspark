package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Conn struct {
	db *pgxpool.Pool
}

func NewConn() (*Conn, error) {
	const (
		host     = "localhost"
		port     = "5433"
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// ParseConfig takes the connection string to generate a config
	config, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		return nil, err
	}

	// MinConns is the minimum number of connections kept open by the pool.
	// The pool will not proactively create this many connections, but once this many have been established,
	// it will not close idle connections unless the total number exceeds MaxConns.
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	// MaxConns is the maximum number of connections that can be opened to PostgreSQL.
	// This limit can be used to prevent overwhelming the PostgreSQL server with too many concurrent connections.
	config.MaxConns = 30

	config.HealthCheckPeriod = 5 * time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Conn{db: db}, nil
}

func (c *Conn) InsertBook(ctx context.Context, newBook NewBook) (Book, error) {

	query := `
		INSERT INTO books (
		                   title, author_name,author_email, 
		                   description, category, 
		                   price, stock
		                  
		                   )
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id int
	err := c.db.QueryRow(
		ctx, query, newBook.Title, newBook.AuthorName,
		newBook.AuthorEmail, newBook.Description, newBook.Category,
		newBook.Price, newBook.Stock,
	).Scan(&id)

	if err != nil {
		//log.Println(err)
		return Book{}, fmt.Errorf("unable to insert book: %w", err)
	}

	b := Book{
		ID:          id,
		Title:       newBook.Title,
		AuthorName:  newBook.AuthorName,
		AuthorEmail: newBook.AuthorEmail,
		Description: newBook.Description,
		Category:    newBook.Category,
		Price:       newBook.Price,
		Stock:       newBook.Stock,
	}
	return b, nil

}

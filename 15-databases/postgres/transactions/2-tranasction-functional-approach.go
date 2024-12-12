package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

var DB *sql.DB

func init() {
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
	var err error
	DB, err = sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	err := DB.Ping()
	if err != nil {
		panic(err)
	}

	err = Update()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Done")
}

func Update() error {
	f := func(tx *sql.Tx) error {
		updateQuery := `UPDATE author
					SET name = $1
					WHERE email = $2;`

		_, err := tx.Exec(updateQuery, "ABC", "john.doe@example.com")
		if err != nil {
			return fmt.Errorf("update error: %w", err)
		}

		_, err = tx.Exec(updateQuery, "John1", "john.doe@example.com")
		if err != nil {
			log.Println(err)
			return fmt.Errorf("update error: %w", err)
		}
		return nil
	}

	err := withTx(context.Background(), f)
	if err != nil {
		return err
	}
	return nil
}

func withTx(ctx context.Context, fn func(*sql.Tx) error) error {

	tx, err := DB.BeginTx(ctx, nil)
	err = fn(tx)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return fmt.Errorf("rollback error: %w", err)
		}
		return fmt.Errorf("transaction error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit error: %w", err)
	}
	return nil

}

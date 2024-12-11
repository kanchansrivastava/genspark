/*
q2. Write a select (id,name,email) from tableName query to fetch all the records
    Steps:-
        user query method
        run a for loop on rows.Next()
        inside the loop scan values using rows.Scan
        print things inside the loop

*/

package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)


func CreateConnection() (*pgxpool.Pool, error) {

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
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.MaxConns = 30

	config.HealthCheckPeriod = 5 * time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return db, nil
}


func main() {
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	getAllUsers(db)
}


func getAllUsers(db *pgxpool.Pool) {
	query := `select id, name, email from users;`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("Unable to fetch users: %v\n", err)
	}
	for rows.Next() {
		var id int
		var name, email string

		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Printf("Row scan failed: %v\n", err)
			continue
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

}

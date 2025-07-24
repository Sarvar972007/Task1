package main

import (
	"Homework_mini_code-1/app"
	"Homework_mini_code-1/repo"
	context2 "context"
	"database/sql"
	"time"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection parameters
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345678"
		dbname   = "postgres"
	)

	// Create connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Error closing connection:", err)
		}
	}()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")

	//err = repo.GoMigrationUp(db)
	//if err != nil {
	//	log.Fatal("Error adding table:", err)
	//}
	//if err = repo.GoMigrationDown(db); err != nil {
	//	log.Fatal("Error removing table:", err)
	//}

	ctx, cancel := context2.WithTimeout(context2.Background(), 3*time.Second)
	defer cancel()

	storage := &repo.Storage{Db: db}
	userService := &service.UserService{Repo: storage}

	client, err := userService.GetUserProfile(ctx, 6)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", client)
}

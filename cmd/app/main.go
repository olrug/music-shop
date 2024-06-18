package main

import (
	"fmt"
	"log"
	"main/api"
	"main/repo"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	dbConnection := os.Getenv("DB_CONNECTION")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	connStr := fmt.Sprintf("%s://%s:%s@localhost:%s/%s", dbConnection, dbUsername, dbPassword, dbPort, dbDatabase)

	db, err := repo.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.Handle()
	log.Fatal(api.ListenAndServe("localhost:8090"))
}

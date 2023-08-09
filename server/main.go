package main

import (
	"database/sql"
	"github.com/mateusmatinato/client-server-api/server/exchangerate"
	sqlDomain "github.com/mateusmatinato/client-server-api/server/query"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	exchangeRateClient := exchangerate.NewClient()

	db := initializeDB()
	exchangeRateRepository := exchangerate.NewRepository(db)
	exchangeRateHandler := exchangerate.NewHandler(exchangeRateClient, exchangeRateRepository)

	mux := http.NewServeMux()
	mux.Handle("/cotacao", http.HandlerFunc(exchangeRateHandler.GetExchangeRate))

	log.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func initializeDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(sqlDomain.InitializeTable); err != nil {
		panic(err)
	}
	return db
}

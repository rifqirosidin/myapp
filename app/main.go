package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"error", "DB connection failed"})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"error", "Cannot ping DB"})
		return
	}

	json.NewEncoder(w).Encode(Message{"success", "Connected to MySQL"})
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api", handler)
	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

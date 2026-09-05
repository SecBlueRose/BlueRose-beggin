package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Response struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	DB      string `json:"db"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "connected"
	if err := db.Ping(); err != nil {
		dbStatus = "disconnected"
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp := Response{
		Status:  "ok",
		Service: "aegis-core-api",
		DB:      dbStatus,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var err error

	dbHost := getEnv("DB_HOST", "postgres")
	dbUser := getEnv("DB_USER", "aegis_admin")
	dbPass := getEnv("DB_PASSWORD", "hsrhi3rd*")
	dbName := getEnv("DB_NAME", "aegis_db")
	dbPort := getEnv("DB_PORT", "5432")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("DB Configuration error: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Printf("Warning:DB is not available %v", err)
	} else {
		log.Println("PostgreSQL successfully connected!")
	}

	http.HandleFunc("/health", healthHandler)

	log.Println("Aegis Core API is listening on port :8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server running error: %v", err)
	}
}

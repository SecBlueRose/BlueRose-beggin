package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response структура описывает JSON, который мы вернем клиенту
type Response struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// healtHandler обрабатывает запросы на /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Указываем заголовок, что ответ будет в формате JSON
	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Status:  "ok",
		Service: "aegis-core-api",
	}
	// Преобразуем структуру в JSON и отправляем в ответ клиенту
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Aegis Core API is listening on port :8080...")

	// Запускаем HTTP-сервер на порту 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: %v", err)
	}
}


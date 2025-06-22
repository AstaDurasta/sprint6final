package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Сервер запущен на http://localhost:8080")

	err := srv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

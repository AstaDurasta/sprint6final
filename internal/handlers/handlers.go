package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(html)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("поступил get запрос"))
		return
	}
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "не удалось прочитать файл", http.StatusInternalServerError)
		return
	}

	newText := service.MorseOrText(string(content))

	dst, err := os.Create(handler.Filename + "_" + time.Now().UTC().Format("20060102_150405"))
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		fmt.Printf("%s", err)
		return
	}
	defer dst.Close()

	_, err = dst.Write([]byte(newText))
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(newText))
}

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
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = res.Write(html)

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
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
	_, _ = w.Write([]byte(newText))
}

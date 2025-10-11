package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(res http.ResponseWriter, req *http.Request) {
	path := "index.html"

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	http.ServeFile(res, req, path)
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)
	fileName := "myFile"

	file, handler, err := req.FormFile(fileName)
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := os.ReadFile(handler.Filename)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	result := service.Convert(string(data))
	newFileName := time.Now().UTC().String() + filepath.Ext(handler.Filename)

	if err := os.WriteFile(newFileName, []byte(result), 0755); err != nil {
		log.Fatal(err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}

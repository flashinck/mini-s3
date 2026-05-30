package main

import (
	"fmt"
	"io"
	"net/http"
)

var memoryStorage = make(map[string][]byte)

func handleObject(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Параметр name обязателен", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPut {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		memoryStorage[name] = data

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Файл успешно сохранен")
		return
	}

	if r.Method == http.MethodGet {
		data, exists := memoryStorage[name]

		if !exists {
			http.Error(w, "Файл не найден", http.StatusNotFound)
			return
		}

		w.Write(data)
		return
	}

	w.WriteHeader(http.Status.MethodNotAllowed)
}

func main() {
	fmt.Println("Mini-S3 v.0.0.1 запускается... ")
	http.HandleFunc("/object", handleObject)
	fmt.Println("Сервер слушает порт :8080")
	http.ListenAndServe(":8080", nil)
}

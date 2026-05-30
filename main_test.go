package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleObject_Put(t *testing.T) {
	// 1. Создаем виртуальное тело запроса (наш файл)
	fileContent := []byte("контент для автотеста")
	body := bytes.NewBuffer(fileContent)

	// 2. Создаем виртуальный PUT-запрос
	req, err := http.NewRequest(http.MethodPut, "/object?name=testfile.txt", body)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	// 3. Создаем рекордер для записи ответа
	rr := httptest.NewRecorder()

	// 4. Напрямую вызываем наш хендлер, передавая туда виртуальный запрос и рекордер
	handleObject(rr, req)

	// 5. ПРОВЕРКА: Ожидаем статус 201 Created
	if rr.Code != http.StatusCreated {
		t.Errorf("Хендлер вернул неправильный статус: получили %d, ожидали %d", rr.Code, http.StatusCreated)
	}

	// 6. ПРОВЕРКА: Убеждаемся, что данные действительно записались в нашу карту memoryStorage
	storedData, exists := memoryStorage["testfile.txt"]
	if !exists {
		t.Errorf("Файл не был сохранен в memoryStorage")
	}
	if string(storedData) != string(fileContent) {
		t.Errorf("Данные в памяти не совпадают: получили %s, ожидали %s", string(storedData), string(fileContent))
	}
}

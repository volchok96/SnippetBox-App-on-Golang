package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello volchok96"))
}

func main() {
	// Инициализация нового рутера
	// Функция "home" как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	// Два параметра: TCP-адрес сети для прослушивания и созданный роутер
	// Функция log.Fatal() для логирования ошибок
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

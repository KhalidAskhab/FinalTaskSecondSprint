package main

import (
	"awesomeProject/internal"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", internal.CalculateHandler)
	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/orchestrator"
)

func main() {
	repo := orchestrator.NewRepository()
	service := orchestrator.NewService(repo)
	handler := orchestrator.NewHandler(service)

	http.HandleFunc("/api/v1/calculate", handler.Calculate)
	http.HandleFunc("/api/v1/expressions", handler.GetExpressions)
	http.HandleFunc("/api/v1/expressions/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/v1/expressions/"):] // Извлекаем ID из URL
		handler.GetExpressionByID(w, r, id)            // Передаем ID в метод
	})
	http.HandleFunc("/internal/task", handler.GetTask)
	http.HandleFunc("/internal/task/result", handler.PostTaskResult)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

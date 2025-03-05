package main

import (
	"fmt"
	"net/http"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/agent"
	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/orchestrator"
)

func startOrchestrator() {
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

	fmt.Println("Orchestrator is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting orchestrator:", err)
	}
}

func startAgent() {
	fmt.Println("Agent is starting...")
	agent.StartAgent()
}

func main() {
	// Запускаем оркестратор в отдельной горутине
	
	go startOrchestrator()

	// Запускаем агента
	startAgent()

	// Бесконечный цикл, чтобы main не завершился
	select {}
}

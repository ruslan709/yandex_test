package agent

import (
	"log"
	"time"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/models"
)

type Worker struct {
	client *Client
}

func NewWorker(client *Client) *Worker {
	return &Worker{client: client}
}

func (w *Worker) Start() {
	for {
		task, err := w.client.GetTask()
		if err != nil {
			log.Println("No tasks available, waiting...")
			time.Sleep(1 * time.Second)
			continue
		}

		if task == nil { // Проверяем, что task не nil
			log.Println("Received nil task, skipping...")
			time.Sleep(1 * time.Second)
			continue
		}

		result := w.compute(task)
		if err := w.client.PostTaskResult(task.ID, result); err != nil {
			log.Println("Failed to post task result:", err)
		}
	}
}

func (w *Worker) compute(task *models.Task) float64 {
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

	switch task.Operation {
	case "addition":
		return task.Arg1 + task.Arg2
	case "subtraction":
		return task.Arg1 - task.Arg2
	case "multiplication":
		return task.Arg1 * task.Arg2
	case "division":
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}

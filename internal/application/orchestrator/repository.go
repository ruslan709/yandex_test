package orchestrator

import (
	"errors"
	"sync"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/models"
)

type Repository struct {
	expressions map[int]*models.Expression
	tasks       map[int]*models.Task
	nextID      int
	mu          sync.Mutex
}

func NewRepository() *Repository {
	return &Repository{
		expressions: make(map[int]*models.Expression),
		tasks:       make(map[int]*models.Task),
		nextID:      1,
	}
}

func (r *Repository) GetNextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++
	return id
}

func (r *Repository) AddExpression(id int, expression string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.expressions[id] = &models.Expression{
		ID:     id,
		Status: "pending",
		Result: 0,
	}
}

func (r *Repository) GetExpressions() []models.Expression {
	r.mu.Lock()
	defer r.mu.Unlock()

	expressions := make([]models.Expression, 0, len(r.expressions))
	for _, expr := range r.expressions {
		expressions = append(expressions, *expr)
	}
	return expressions
}

func (r *Repository) GetExpressionByID(id int) (*models.Expression, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	expr, ok := r.expressions[id]
	if !ok {
		return nil, errors.New("expression not found")
	}
	return expr, nil
}

func (r *Repository) GetTask() (*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, task := range r.tasks {
		if task.Status == "pending" {
			task.Status = "in_progress"
			return task, nil
		}
	}
	return nil, errors.New("no tasks available")
}

func (r *Repository) PostTaskResult(id int, result float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	task.Result = result
	task.Status = "completed"
	return nil
}

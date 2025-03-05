package orchestrator

import (
	"sync"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/models"
)

type Service struct {
	repo *Repository
	mu   sync.Mutex
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddExpression(expression string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.repo.GetNextID()
	s.repo.AddExpression(id, expression)
	return id, nil
}

func (s *Service) GetExpressions() []models.Expression {
	return s.repo.GetExpressions()
}

func (s *Service) GetExpressionByID(id int) (*models.Expression, error) {
	return s.repo.GetExpressionByID(id)
}

func (s *Service) GetTask() (*models.Task, error) {
	return s.repo.GetTask()
}

func (s *Service) PostTaskResult(id int, result float64) error {
	return s.repo.PostTaskResult(id, result)
}

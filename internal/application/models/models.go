package models

type Expression struct {
	ID     int     `json:"id"`         // Уникальный идентификатор выражения
	Status string  `json:"status"`     // Статус вычисления (например, "pending", "in_progress", "completed")
	Result float64 `json:"result"`     // Результат вычисления (если вычисление завершено)
	Expr   string  `json:"expression"` // Само выражение (например, "2 + 2 * 2")
}

type Task struct {
	ID            int     `json:"id"`             // Уникальный идентификатор задачи
	Arg1          float64 `json:"arg1"`           // Первый аргумент
	Arg2          float64 `json:"arg2"`           // Второй аргумент
	Operation     string  `json:"operation"`      // Операция (например, "addition", "subtraction", "multiplication", "division")
	OperationTime int     `json:"operation_time"` // Время выполнения операции в миллисекундах
	Status        string  `json:"status"`         // Статус задачи (например, "pending", "in_progress", "completed")
	Result        float64 `json:"result"`         // Результат выполнения задачи
}

const (
	OperationAddition       = "addition"
	OperationSubtraction    = "subtraction"
	OperationMultiplication = "multiplication"
	OperationDivision       = "division"
)

const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)

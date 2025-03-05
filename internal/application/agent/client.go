package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/models"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func (c *Client) GetTask() (*models.Task, error) {
	resp, err := http.Get(c.baseURL + "/internal/task")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get task: status code %d", resp.StatusCode)
	}

	var result struct {
		Task *models.Task `json:"task"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Task == nil { // Проверяем, что задача не nil
		return nil, fmt.Errorf("no task available")
	}

	return result.Task, nil
}

func (c *Client) PostTaskResult(id int, result float64) error {
	reqBody, err := json.Marshal(map[string]interface{}{
		"id":     id,
		"result": result,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(c.baseURL+"/internal/task/result", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

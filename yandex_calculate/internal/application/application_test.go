package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockCalc(expression string) (float64, error) {
	switch expression {
	case "2+2":
		return 4, nil
	case "2/0":
		return 0, fmt.Errorf("division by zero")
	case "2+2a":
		return 0, fmt.Errorf("invalid character")
	case "unknown":
		return 0, fmt.Errorf("calculation error")
	default:
		return 0, nil
	}
}

func TestCalculateHandler(t *testing.T) {
	calcFunc = mockCalc

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   Response
	}{
		{
			name:           "Valid expression",
			requestBody:    `{"expression":"2+2"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Result: "4"},
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"expression":2+2}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Response{Error: "Invalid request payload"},
		},
		{
			name:           "Invalid characters in expression",
			requestBody:    `{"expression":"2+2a"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Response{Error: "Invalid characters in expression"},
		},
		{
			name:           "Division by zero",
			requestBody:    `{"expression":"2/0"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Response{Error: "Division by zero error"},
		},
		{
			name:           "Calculation error",
			requestBody:    `{"expression":"unknown"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Response{Error: "Calculation error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(calculateHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var resBody Response
			err = json.NewDecoder(rr.Body).Decode(&resBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedBody, resBody)
		})
	}
}

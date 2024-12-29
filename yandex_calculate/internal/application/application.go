package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/pkg/calculation"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}
var calcFunc = calculation.Calc
func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		res.Error = "Invalid request payload"
		json.NewEncoder(w).Encode(res)
		return
	}

	validExpression := regexp.MustCompile(`^[0-9+\-*/\s()]+$`)
	if !validExpression.MatchString(req.Expression) || strings.TrimSpace(req.Expression) == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		res.Error = "Invalid characters in expression"
		json.NewEncoder(w).Encode(res)
		return
	}

	if strings.Contains(req.Expression, "/0") {
		w.WriteHeader(http.StatusUnprocessableEntity)
		res.Error = "Division by zero error"
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		log.Println("calculation failed with error: ", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		res.Error = "Calculation error"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = fmt.Sprintf("%v", result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func RunServer() {
	http.Handle("/api/v1/calculate", PanicMiddleware(http.HandlerFunc(calculateHandler)))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

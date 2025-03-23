package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yingtu35/GoCalcAPI/internal/model"
	"github.com/yingtu35/GoCalcAPI/pkg/calculator"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("OK"))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, calculator.Add)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, calculator.Subtract)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, calculator.Multiply)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	operationHandler(w, r, calculator.Divide)
}

func operationHandler(w http.ResponseWriter, r *http.Request, operation func(int, int) (int, error)) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	input, err := parseInput(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input JSON")
		return
	}
	if input.Value1 == nil || input.Value2 == nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input: value1 and value2 are required")
		return
	}
	result, err := operation(*(input.Value1), *(input.Value2))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithResult(w, result)
}

func parseInput(r *http.Request) (*model.Input, error) {
	body := r.Body
	defer body.Close()

	// Decode the input into the model.Input struct
	var input model.Input
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}

func respondWithError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.Output{Error: errMsg})
}

func respondWithResult(w http.ResponseWriter, result int) {
	output := model.Output{Result: &result}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Fatal(err)
	}
}

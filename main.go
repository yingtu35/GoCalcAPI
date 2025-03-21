package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Input struct {
	Value1 *int `json:"value1"`
	Value2 *int `json:"value2"`
}

type Output struct {
	Result *int   `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()

	// Decode the input into the Input struct
	var input Input
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input JSON"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}
	// Input validation
	if input.Value1 == nil || input.Value2 == nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input: value1 and value2 are required"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}

	result := *(input.Value1) + *(input.Value2)
	output := Output{Result: &result}

	// Encode the output into JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Fatal(err)
	}
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()

	// Decode the input into the Input struct
	var input Input
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input JSON"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}
	// Input validation
	if input.Value1 == nil || input.Value2 == nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input: value1 and value2 are required"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}

	result := *(input.Value1) - *(input.Value2)
	output := Output{Result: &result}

	// Encode the output into JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Fatal(err)
	}
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()

	// Decode the input into the Input struct
	var input Input
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input JSON"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}
	// Input validation
	if input.Value1 == nil || input.Value2 == nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input: value1 and value2 are required"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}

	result := *(input.Value1) * *(input.Value2)
	output := Output{Result: &result}

	// Encode the output into JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Fatal(err)
	}
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()

	// Decode the input into the Input struct
	var input Input
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input JSON"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}
	// Input validation
	if input.Value1 == nil || input.Value2 == nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Invalid input: value1 and value2 are required"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}

	// Division by zero check
	if *(input.Value2) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := "Cannot divide by zero"
		json.NewEncoder(w).Encode(Output{Error: errMsg})
		return
	}

	result := *(input.Value1) / *(input.Value2)
	output := Output{Result: &result}

	// Encode the output into JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/subtract", subtractHandler)
	mux.HandleFunc("/multiply", multiplyHandler)
	mux.HandleFunc("/divide", divideHandler)

	wrappedMux := NewLogger(mux)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}

package model

type Input struct {
	Value1 *int `json:"value1"`
	Value2 *int `json:"value2"`
}

type Output struct {
	Result *int   `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

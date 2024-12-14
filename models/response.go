package models

// Response представляет ответ от LLama API
type Response struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

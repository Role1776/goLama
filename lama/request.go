package lama

type createPayload struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

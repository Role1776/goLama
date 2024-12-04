package lama

func createPayload(model, prompt string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"prompt": prompt,
	}
}

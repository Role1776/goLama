package lama

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/Role1776/goLama/models"
)

func GenerateResponse(url string, model string, prompt string, timeout time.Duration) (<-chan models.Response, <-chan error) {
	payload := createPayload{
		Model:  model,
		Prompt: prompt,
	}

	payloadJSON, err := json.Marshal(payload)

	errChan := make(chan error)

	if err != nil {
		errChan <- fmt.Errorf("failed to marshal payload: %v", err)

		return nil, errChan
	}

	respChan := make(chan models.Response)

	go func() {
		defer close(respChan)
		defer close(errChan)

		resp, err := sendRequest(url, payloadJSON, timeout)
		if err != nil {
			errChan <- fmt.Errorf("failed to send request: %v", err)

			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errChan <- fmt.Errorf("unexpected status code %d: %v", resp.StatusCode, fmt.Errorf("expected 200 OK"))

			return
		}

		handleResponse(resp.Body, respChan, errChan)

	}()

	return respChan, errChan
}

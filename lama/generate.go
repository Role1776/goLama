package lama

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Role1776/goLama/utils"
)

func GenerateResponse(url, model, prompt string) (string, error) {
	payload := createPayload(model, prompt)

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("failed to marshal payload for model %s and prompt %s: %v", model, prompt, err)
		utils.Logger(err)
		return "", err
	}

	resp, err := sendRequest(url, payloadJSON) 
	if err != nil {
		err = fmt.Errorf("failed to send request to URL %s: %v", url, err)
		utils.Logger(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code %d from URL %s: expected 200 OK", resp.StatusCode, url)
		utils.Logger(err)
		return "", err
	}

	fullResponse, err := handleResponse(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to handle response body from URL %s: %v", url, err)
		utils.Logger(err)
		return "", err
	}

	return fullResponse, nil
}


package lama

import (
	"encoding/json"
	"io"

	"github.com/Role1776/goLama/models"
)

func handleResponse(respBody io.Reader) (string, error) {
	var fullResponse string
	decoder := json.NewDecoder(respBody)

	for {
		var response models.Response
		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		fullResponse += response.Response

		if response.Done {
			break
		}
	}

	return fullResponse, nil
}
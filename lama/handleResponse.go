package lama

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Role1776/goLama/models"
)

func handleResponse(body io.Reader, respChan chan<- models.Response, errChan chan<- error) {
	decoder := json.NewDecoder(body)
	for {
		var resp models.Response
		if err := decoder.Decode(&resp); err != nil {
			errChan <- fmt.Errorf("failed to decode response: %v", err)
			return
		}

		respChan <- resp

		if resp.Done {
			return
		}
	}
}

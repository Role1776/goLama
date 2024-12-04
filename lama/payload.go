package lama

import (
	"bytes"
	"net/http"
	"time"
)

const (
	defaultTimeout = 100 * time.Second
)

func sendRequest(url string, payloadJSON []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	
	
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

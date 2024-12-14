package lama

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Role1776/goLama/models"
)

func createTestServer(handler http.HandlerFunc) *httptest.Server {
	server := httptest.NewServer(handler)
	return server
}

func TestGenerateResponse_Success(t *testing.T) {
	expectedResponse := []models.Response{
		{Response: "Part 1", Done: false},
		{Response: "Part 2", Done: true},
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var payload createPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("failed to decode payload: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if payload.Model != "testModel" || payload.Prompt != "testPrompt" {
			t.Errorf("Unexpected payload: got model: %s, prompt: %s", payload.Model, payload.Prompt)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		for _, resp := range expectedResponse {
			if err := enc.Encode(resp); err != nil {
				t.Errorf("failed to encode: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})

	server := createTestServer(handler)
	defer server.Close()

	respChan, errChan := GenerateResponse(server.URL, "testModel", "testPrompt", time.Second)

	var responses []models.Response
	done := false
	for !done {
		select {
		case resp, ok := <-respChan:
			if !ok {
				done = true
				continue
			}
			responses = append(responses, resp)
			if resp.Done {
				done = true
			}
		case err := <-errChan:
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(responses) != len(expectedResponse) {
		t.Errorf("Expected %d responses, got %d", len(expectedResponse), len(responses))
	}

	for i, resp := range responses {
		if resp != expectedResponse[i] {
			t.Errorf("Response at index %d does not match", i)
		}
	}
}

func TestGenerateResponse_RequestError(t *testing.T) {
	respChan, errChan := GenerateResponse("invalid-url", "testModel", "testPrompt", time.Second)

	select {
	case resp := <-respChan:
		t.Fatalf("Expected error, but got response: %+v", resp)
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}
	}
}

func TestGenerateResponse_BadStatusCode(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	server := createTestServer(handler)
	defer server.Close()

	respChan, errChan := GenerateResponse(server.URL, "testModel", "testPrompt", time.Second)

	select {
	case resp := <-respChan:
		t.Fatalf("Expected error, but got response: %+v", resp)
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}
		if !strings.Contains(err.Error(), "unexpected status code 400") {
			t.Errorf("Error message should contain 'unexpected status code 400', but got '%s'", err.Error())
		}
	}
}

func TestHandleResponse_Success(t *testing.T) {
	expectedResponse := []models.Response{
		{Response: "Part 1", Done: false},
		{Response: "Part 2", Done: true},
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for _, resp := range expectedResponse {
		if err := enc.Encode(resp); err != nil {
			t.Fatalf("failed to encode test data: %v", err)
		}
	}

	respChan := make(chan models.Response)
	errChan := make(chan error)
	go handleResponse(&buf, respChan, errChan)

	var responses []models.Response
	done := false
	for !done {
		select {
		case resp, ok := <-respChan:
			if !ok {
				done = true
				continue
			}
			responses = append(responses, resp)
			if resp.Done {
				done = true
			}
		case err := <-errChan:
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(responses) != len(expectedResponse) {
		t.Errorf("Expected %d responses, got %d", len(expectedResponse), len(responses))
	}

	for i, resp := range responses {
		if resp != expectedResponse[i] {
			t.Errorf("Response at index %d does not match", i)
		}
	}
}

func TestHandleResponse_DecodeError(t *testing.T) {
	invalidJSON := bytes.NewBufferString(`invalid json`)
	respChan := make(chan models.Response)
	errChan := make(chan error)

	go handleResponse(invalidJSON, respChan, errChan)

	select {
	case resp := <-respChan:
		t.Fatalf("Expected error, but got response: %+v", resp)
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}
		if !strings.Contains(err.Error(), "failed to decode response") {
			t.Errorf("Error message should contain 'failed to decode response', but got '%s'", err.Error())
		}
	}
}

func TestSendRequest_Success(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type header should be application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	server := createTestServer(handler)
	defer server.Close()

	payload := []byte(`{"key":"value"}`)
	resp, err := sendRequest(server.URL, payload, 1*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}
func TestSendRequest_Timeout(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	})

	server := createTestServer(handler)
	defer server.Close()
	payload := []byte(`{"key":"value"}`)

	_, err := sendRequest(server.URL, payload, 1*time.Second)

	if err == nil {
		t.Fatalf("Expected timeout error, but got nil")
	}

	var netError interface {
		Timeout() bool
		Temporary() bool
	}

	if !errors.As(err, &netError) || !netError.Timeout() {
		t.Errorf("Expected timeout error, got: %v", err)
	}
}

func TestSendRequest_RequestError(t *testing.T) {
	payload := []byte(`{"key":"value"}`)
	_, err := sendRequest("invalid-url", payload, time.Second)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

}

func TestSyncResponse_Success(t *testing.T) {
	expectedResponse := []models.Response{
		{Response: "Part 1", Done: false},
		{Response: "Part 2", Done: true},
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		for _, resp := range expectedResponse {
			if err := enc.Encode(resp); err != nil {
				t.Fatalf("failed to encode: %v", err)
			}
		}
	})

	server := createTestServer(handler)
	defer server.Close()

	result := SyncResponse(server.URL, "testModel", "testPrompt", time.Second)
	expectedResult := "Part 1Part 2"

	if result != expectedResult {
		t.Errorf("Expected result '%s', got '%s'", expectedResult, result)
	}
}

func TestSyncResponse_Error(t *testing.T) {
	result := SyncResponse("invalid-url", "testModel", "testPrompt", time.Second)
	if result != "" {
		t.Fatalf("Expected empty string but got: %s", result)
	}
}

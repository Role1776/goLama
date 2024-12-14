package lama

import (
	"log"
	"strings"
	"time"
)

func SyncResponse(url string, model string, prompt string, timeout time.Duration) string {
	resp, errChan := GenerateResponse(url, model, prompt, timeout)
	var allResp strings.Builder
	done := false
	for !done {

		select {

		case resp, ok := <-resp:
			if !ok {
				done = true
				continue
			}

			allResp.WriteString(resp.Response)
			done = resp.Done

		case err := <-errChan:
			if err != nil {
				log.Println("Error:", err)
			}
			done = true
		}

	}
	return allResp.String()
}

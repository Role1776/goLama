package utils

import (
	"log"
)

func Logger(err error) {
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
}
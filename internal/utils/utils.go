package utils

import (
	"log"
	"os"
)

func WorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current dir: %v", err)
	}
	return dir
}

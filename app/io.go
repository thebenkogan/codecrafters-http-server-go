package main

import (
	"io"
	"os"
)

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil
	} else {
		defer file.Close()
		buffer, _ := io.ReadAll(file)
		return buffer
	}
}

func WriteFile(path string, buffer []byte) {
	os.WriteFile(path, buffer, 0644)
}

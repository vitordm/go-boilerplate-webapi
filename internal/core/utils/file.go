package utils

import (
	"errors"
	"io"
	"os"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func ReadAllFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// Leia todo o conteúdo do arquivo
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

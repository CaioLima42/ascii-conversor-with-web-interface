package utils

import (
	"os"
	"path/filepath"
	"fmt"
)

func ReadTemplate(fileName string) ([]byte, error) {
	path := filepath.Join("web", "templates", fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler template: %v", err)
	}
	return data, nil
}

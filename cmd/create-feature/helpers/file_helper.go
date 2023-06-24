package helpers

import (
	"os"
)

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ReadFileAsString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	return string(data), err
}

func WriteFileFromString(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}

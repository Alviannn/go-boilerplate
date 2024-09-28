package helpers

import (
	"os"
	"path"
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

func ReadCompleteFilesInDir(filePath string) (filePaths []string, err error) {
	fileList, err := os.ReadDir(filePath)
	if err != nil {
		return
	}

	for _, file := range fileList {
		completeFilePath := path.Join(filePath, file.Name())
		if !file.IsDir() {
			filePaths = append(filePaths, completeFilePath)
			continue
		}

		subFilePaths, err := ReadCompleteFilesInDir(completeFilePath)
		if err != nil {
			return nil, err
		}
		filePaths = append(filePaths, subFilePaths...)
	}

	return
}

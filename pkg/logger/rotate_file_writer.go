package logger

import (
	"os"
	"sync"
)

type (
	RotateFileWriter struct {
		currentFilePath  string
		nextFilePathFunc NextFilePathFunc
		writer           *os.File
		mutex            sync.Mutex
	}

	NextFilePathFunc func() string
)

func NewRotateFileWriter(nextFunc NextFilePathFunc) *RotateFileWriter {
	return &RotateFileWriter{
		nextFilePathFunc: nextFunc,
	}
}

func (w *RotateFileWriter) rotateFile(nextFilePath string) (err error) {
	// Closes previous file writer if there were any.
	if err = w.Close(); err != nil {
		return
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	fileWriter, err := os.OpenFile(nextFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}

	w.writer = fileWriter
	w.currentFilePath = nextFilePath
	return
}

func (w *RotateFileWriter) Write(p []byte) (n int, err error) {
	if nextFilePath := w.nextFilePathFunc(); w.currentFilePath != nextFilePath {
		if err = w.rotateFile(nextFilePath); err != nil {
			return
		}
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.writer.Write(p)
}

func (w *RotateFileWriter) Close() (err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.writer == nil {
		return
	}

	err = w.writer.Close()
	if err == nil {
		w.writer = nil
	}
	return
}

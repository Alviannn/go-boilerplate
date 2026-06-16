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

	PerformRotateFileParam struct {
		NextFilePath string
		IsUseLock    bool
	}
)

func NewRotateFileWriter(nextFunc NextFilePathFunc) *RotateFileWriter {
	return &RotateFileWriter{
		nextFilePathFunc: nextFunc,
	}
}

func (w *RotateFileWriter) Write(p []byte) (n int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	nextFilePath := w.nextFilePathFunc()

	if w.writer == nil || w.currentFilePath != nextFilePath {
		err = w.performRotateFile(PerformRotateFileParam{
			NextFilePath: nextFilePath,
			IsUseLock:    false, // already locked
		})
		if err != nil {
			return
		}
	}

	return w.writer.Write(p)
}

func (w *RotateFileWriter) performRotateFile(param PerformRotateFileParam) (err error) {
	// Closes previous file writer if there were any.
	if err = w.performClose(param.IsUseLock); err != nil {
		return
	}

	var (
		defaultFilePerm = os.FileMode(0644)
		defaultFileFlag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	)

	fileWriter, err := os.OpenFile(param.NextFilePath, defaultFileFlag, defaultFilePerm)
	if err != nil {
		return
	}

	w.writer = fileWriter
	w.currentFilePath = param.NextFilePath
	return
}

func (w *RotateFileWriter) performClose(isUseLock bool) (err error) {
	if isUseLock {
		w.mutex.Lock()
		defer w.mutex.Unlock()
	}

	if w.writer == nil {
		return
	}

	err = w.writer.Close()
	if err == nil {
		w.writer = nil
		w.currentFilePath = ""
	}
	return
}

func (w *RotateFileWriter) Close() (err error) {
	return w.performClose(true)
}

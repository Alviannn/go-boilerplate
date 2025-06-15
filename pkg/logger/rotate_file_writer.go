package logger

import (
	"os"
)

type (
	RotateFileWriter struct {
		currentFilePath  string
		nextFilePathFunc NextFilePathFunc
		writer           *os.File
	}

	NextFilePathFunc func() string
)

func NewRotateFileWriter(nextFunc NextFilePathFunc) (writer *RotateFileWriter, err error) {
	writer = &RotateFileWriter{
		nextFilePathFunc: nextFunc,
	}
	err = writer.rotateFile(writer.nextFilePathFunc())
	return
}

func (w *RotateFileWriter) rotateFile(nextFilePath string) (err error) {
	// Closes previous file writer if there were any.
	if err = w.Close(); err != nil {
		return
	}

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
	return w.writer.Write(p)
}

func (w *RotateFileWriter) Close() (err error) {
	if w.writer == nil {
		return
	}

	err = w.writer.Close()
	if err == nil {
		w.writer = nil
	}
	return
}

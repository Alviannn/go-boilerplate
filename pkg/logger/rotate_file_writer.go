package logger

import (
	"os"
	"strings"
	"time"
)

type RotateFileWriter struct {
	filenameWithFormat string
	currentFilePath    string

	fileWriter *os.File
}

func NewRotateFileWriter(filenameWithFormat string) (writer *RotateFileWriter, err error) {
	writer = &RotateFileWriter{
		filenameWithFormat: filenameWithFormat,
	}

	err = writer.rotateFile()
	return
}

func (w *RotateFileWriter) getExpectedFilePath() string {
	currentDate := time.Now().Format(time.DateOnly)
	return strings.ReplaceAll(w.filenameWithFormat, "{date}", currentDate)
}

func (w *RotateFileWriter) rotateFile() (err error) {
	// Close the old (or previous) file writer if there were any.
	if err = w.Close(); err != nil {
		return
	}

	expectedFilePath := w.getExpectedFilePath()
	fileWriter, err := os.OpenFile(expectedFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	w.fileWriter = fileWriter
	return
}

func (w *RotateFileWriter) Write(p []byte) (n int, err error) {
	if w.currentFilePath != w.getExpectedFilePath() {
		if err = w.rotateFile(); err != nil {
			return
		}
	}

	return w.fileWriter.Write(p)
}

func (w *RotateFileWriter) Close() (err error) {
	if w.fileWriter == nil {
		return
	}

	if err = w.fileWriter.Close(); err != nil {
		w.fileWriter = nil
	}

	return err
}

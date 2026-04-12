package logger_test

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/logger"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
)

type (
	loggerTest struct {
		suite.Suite
	}
)

func TestLogger(t *testing.T) {
	suite.Run(t, new(loggerTest))
}

func (s *loggerTest) setupTest() (finishFunc func()) {
	err := configs.LoadWithConfig(configs.LoadParam{IsMock: true})
	if err != nil {
		s.T().Fatal(err)
	}

	configMocked := configs.GetMocked()
	configMocked.LogsDir = "./logs-test"

	return func() {
		os.RemoveAll(configMocked.LogsDir)
	}
}

func (s *loggerTest) prepareLogFile(t *testing.T, fileName string) (writer *logger.RotateFileWriter, err error) {
	t.Helper()

	err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		return
	}

	writer = logger.NewRotateFileWriter(func() string {
		return fileName
	})
	return
}

func (s *loggerTest) TestSetup() {
	finishFunc := s.setupTest()
	defer finishFunc()

	logFilePath := filepath.Join(configs.Default().LogsDir, "test.log")
	logFileWriter, err := s.prepareLogFile(s.T(), logFilePath)

	if logFileWriter != nil {
		defer logFileWriter.Close()

		logger.Setup(logger.SetupParam{
			ConsoleWriter: nil,
			ExtraWriters:  []io.Writer{logFileWriter},
		})
	}

	s.NoError(err)
	s.NotNil(logFileWriter)

	log.Info().Msg("hello world")

	s.NoError(err)
	s.FileExists(logFilePath, "logs file should exist")
}

func (s *loggerTest) TestWriteLog() {
	finishFunc := s.setupTest()
	defer finishFunc()

	logFilePath := filepath.Join(configs.Default().LogsDir, "test.log")
	logFileWriter, err := s.prepareLogFile(s.T(), logFilePath)

	if logFileWriter != nil {
		defer logFileWriter.Close()

		logger.Setup(logger.SetupParam{
			ConsoleWriter: nil,
			ExtraWriters:  []io.Writer{logFileWriter},
		})
	}

	s.Require().NoError(err)

	log.Info().Msg("test info")
	s.FileExists(logFilePath, "logs file should exist")

	logContent, err := os.ReadFile(logFilePath)
	s.NoError(err, "log file should be readable")
	s.NotEmpty(logContent, "log file should not be empty")

	logContentStr := string(logContent)
	s.Contains(logContentStr, "\"level\":\"info\"", "log should contain level info")
	s.Contains(logContentStr, "\"message\":\"test info\"", "log should contain message test info")
}

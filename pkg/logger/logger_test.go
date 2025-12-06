package logger_test

import (
	"fmt"
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/logger"
	"os"
	"path"
	"testing"
	"time"

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

func (s *loggerTest) TestSetup() {
	finishFunc := s.setupTest()
	defer finishFunc()

	var (
		logsDir     = configs.Default().LogsDir
		logFilePath = path.Join(logsDir, fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly)))
		err         = logger.Setup()
	)

	log.Info().Msg("hello world")

	s.NoError(err)
	s.FileExists(logFilePath, "logs file should exist")
}

func (s *loggerTest) TestWriteLog() {
	finishFunc := s.setupTest()
	defer finishFunc()

	var (
		logsDir     = configs.Default().LogsDir
		fileName    = fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
		logFilePath = path.Join(logsDir, fileName)
	)

	err := logger.Setup()
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

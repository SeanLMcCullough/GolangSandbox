package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger provides a configured Logrus Logger
func NewLogger() *logrus.Logger {
	var log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log
}

package logrus

import (
	"os"

	"github.com/kialkuz/task-manager/pkg/logger"
	log "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	entry *log.Entry
}

func NewLogger() logger.LoggerInterface {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(log.WarnLevel)

	return &LogrusLogger{
		entry: log.NewEntry(logger),
	}
}

func (l *LogrusLogger) Info(err error) {
	l.entry.Info(err.Error())
}

func (l *LogrusLogger) Debug(err error) {
	l.entry.Debug(err.Error())
}

func (l *LogrusLogger) Error(err error) {
	l.entry.Error(err.Error())
}

func (l *LogrusLogger) WithFields(fields logger.Field) logger.LoggerInterface {
	if len(fields) == 0 {
		return &LogrusLogger{}
	}

	combined := log.Fields{}
	for fieldName, fieldValue := range fields {
		combined[fieldName] = fieldValue
	}

	return &LogrusLogger{
		entry: l.entry.WithFields(log.Fields(fields)),
	}
}

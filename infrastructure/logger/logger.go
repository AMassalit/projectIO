package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(level string) *Logger {
	log := logrus.New()
	var lev logrus.Level

	switch strings.ToLower(level) {
	case "error":
		lev = 2
	case "warn":
		lev = 3
	case "info":
		lev = 4
	case "debug":
		lev = 5
	default:
		lev = 4
	}

	jsonFormatter := &logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	log.SetLevel(lev)
	log.SetFormatter(jsonFormatter)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	return &Logger{
		log,
	}
}
func (l *Logger) MapFields(fields map[string]interface{}) logrus.Fields {
	log := logrus.Fields{}

	for k, v := range fields {
		log[k] = v
	}

	return log
}

package helper

import "github.com/sirupsen/logrus"

const (
	LoggerLevelTrace = "LoggerLevelTrace"
	LoggerLevelDebug = "LoggerLevelDebug"
	LoggerLevelInfo  = "LoggerLevelInfo"
	LoggerLevelWarn  = "LoggerLeveWarn"
	LoggerLevelError = "LoggerLevelError"
	LoggerLevelFatal = "LoggerLevelFatal"
	LoggerLevelPanic = "LoggerLevelPanic"
)

func Logger(filepath, level, message string) {
	if filepath == "" || level == "" || message == "" {
		logrus.WithFields(
			logrus.Fields{
				"file": "internal/helper/logger.go",
			},
		).Error("All params is required")
	}

	logging := logrus.WithFields(
		logrus.Fields{
			"file": filepath,
		})

	switch level {
	case LoggerLevelDebug:
		logging.Debug(message)
	case LoggerLevelInfo:
		logging.Info(message)
	case LoggerLevelWarn:
		logging.Warn(message)
	case LoggerLevelError:
		logging.Error(message)
	case LoggerLevelFatal:
		logging.Fatal(message)
	case LoggerLevelPanic:
		logging.Panic(message)
	default:
		logging.Error("Level invalid")
	}

}

package logging

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {

	// internal, specific to the underlying logging library
	contextData log.Fields

	// internal, specific to the underlying logging library
	logger log.Logger

	// independant of the underlying logging library
	LogInterface
}

type DataFields map[string]interface{}

type LogInterface interface {
	Debug(message string, data DataFields)
	Info(message string, data DataFields)
	Warn(message string, data DataFields)
	Error(message string, data DataFields)
}

type LogConfig struct {
	AppName    string
	AppVersion string


	// DEBUG, INFO, WARN, ERROR
	Level string
}

func New(config *LogConfig) (Logger, error) {
	// parse through level
	var logLevel log.Level

	switch config.Level {
	case "DEBUG", "debug":
		logLevel = log.DebugLevel
	case "INFO", "info":
		logLevel = log.InfoLevel
	case "WARN", "warn":
		logLevel = log.WarnLevel
	case "ERROR", "error":
		logLevel = log.ErrorLevel
	case "PANIC", "panic":
		fallthrough
	case "FATAL", "fatal":
		fallthrough
	default:
		return Logger{}, fmt.Errorf("Invalid logging level. valid options are DEBUG, INFO, WARN & ERROR")
	}

	// TODO: Add checker for environment

	logger := log.Logger{
		Out:       os.Stdout,
		Formatter: new(log.JSONFormatter),
		Level:     logLevel,
	}

	return Logger{contextData: log.Fields{
		"app-name":    config.AppName,
		"app-version": config.AppVersion,
	}, logger: logger}, nil
}

func (l *Logger) Debug(message string, data ...DataFields) {
	if data == nil {
		l.logger.WithFields(l.contextData).Debug(message)
		return
	}

	newData := l.mapJoin(data[0])
	l.logger.WithFields(newData).Debug(message)
}

func (l *Logger) Info(message string, data ...DataFields) {
	if data == nil {
		l.logger.WithFields(l.contextData).Info(message)
		return
	}

	newData := l.mapJoin(data[0])
	l.logger.WithFields(newData).Info(message)
}

func (l *Logger) Warn(message string, data ...DataFields) {
	if data == nil {
		l.logger.WithFields(l.contextData).Warn(message)
		return
	}

	newData := l.mapJoin(data[0])
	l.logger.WithFields(newData).Warn(message)
}

func (l *Logger) Error(message string, data ...DataFields) {
	if data == nil {
		l.logger.WithFields(l.contextData).Error(message)
		return
	}

	newData := l.mapJoin(data[0])
	l.logger.WithFields(newData).Error(message)
}

// private
func (l *Logger) mapJoin(fieldData DataFields) log.Fields {
	fields := make(log.Fields, len(fieldData)+len(l.contextData))

	for k, v := range fieldData {
		fields[k] = v
	}

	for k, v := range l.contextData {
		fields[k] = v
	}
	return fields
}

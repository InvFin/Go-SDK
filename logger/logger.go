package logger

import (
	"log"
	"os"
)

type Logger struct {
	file     *os.File
	logger   *log.Logger
	logLevel LogLevel
}

type LogLevel int

const (
	InfoLevel LogLevel = iota
	ErrorLevel
	DebugLevel
)

func NewLogger(filename string, logLevel LogLevel) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.Ldate|log.Ltime)

	return &Logger{
		file:     file,
		logger:   logger,
		logLevel: logLevel,
	}, nil
}

func (l *Logger) Info(message string) {
	if l.logLevel <= InfoLevel {
		l.logger.Println("[INFO] " + message)
	}
}

func (l *Logger) Error(message string) {
	if l.logLevel <= ErrorLevel {
		l.logger.Println("[ERROR] " + message)
	}
}

func (l *Logger) Debug(message string) {
	if l.logLevel <= DebugLevel {
		l.logger.Println("[DEBUG] " + message)
	}
}

func (l *Logger) Close() error {
	return l.file.Close()
}

package services

import (
	"log"
	"os"
)

type LoggerService struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewLogger() *LoggerService {
	return &LoggerService{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *LoggerService) Debug(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *LoggerService) Info(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *LoggerService) Warn(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *LoggerService) Error(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}

package helper

import (
	log "github.com/sirupsen/logrus"
)

const (
	// TOPIC for setting topic of log
	TOPIC = "master-service-log"
	// LogTag default log tag
	LogTag = "master-service"
)

type Webhook struct {
	Color  string  `json:"color"`
	Fields []Field `json:"fields"`
	Text   string  `json:"text"`
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

// LogContext function for logging the context of echo
// c string context
// s string scope
func LogContext(c string, s string) *log.Entry {
	return log.WithFields(log.Fields{
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

// Log function for returning entry type
// level log.Level
// message string message of log
// context string context of log
// scope string scope of log
func Log(level log.Level, message string, context string, scope string) {

	entry := LogContext(context, scope)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}

package util

import (
	"log"
	"net/http"
	"os"
	"time"
)

func httpLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Logger logs the standard and error logs to system output according to their levels
var Logger = createLogLevels()

type logLevels struct {
	TRACE *logLevel
	INFO  *logLevel
	ERROR *logLevel
	FATAL *logLevel
}

type logLevel struct {
	level   string
	isError bool
}

func (l *logLevel) Log(message string, err ...error) {
	if l.isError {
		log.SetOutput(os.Stderr)
		log.Printf("[%s] [%s] %s\nError:[%s]", l.level, time.Now(), message, err)
	} else {
		log.SetOutput(os.Stdout)
		log.Printf("[%s] [%s] %s", l.level, time.Now(), message)
	}
}

func createLogLevels() *logLevels {
	trace := &logLevel{"TRACE", false}
	info := &logLevel{"INFO", false}
	err := &logLevel{"ERROR", true}
	fatal := &logLevel{"FATAL", true}

	return &logLevels{
		TRACE: trace,
		INFO:  info,
		ERROR: err,
		FATAL: fatal,
	}
}

//LogMessage specifies the struct which has the parameters to pass to the Log method
type LogMessage struct {
	Message string
	Error   error
}

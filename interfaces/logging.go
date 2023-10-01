package interfaces

import (
	"io"
	"log"
)

type Logger interface {
	Log(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Debug(string, ...any)
}

type PreludeLogger struct {
	writer io.WriteCloser
	prompt struct {
		Info, Warn, Error, Debug string
	}
}

func NewPreludeLogger(w io.WriteCloser) *PreludeLogger {
	if w != nil {
		log.SetOutput(w)
	}
	return &PreludeLogger{
		writer: w,
		prompt: struct {
			Info  string
			Warn  string
			Error string
			Debug string
		}{
			Info:  "[Info]",
			Warn:  "[Warning]",
			Error: "[Error]",
			Debug: "[Debug]",
		},
	}
}

func (logging *PreludeLogger) Log(s string, args ...any) {
	log.Printf(s, args...)
}

func (logging *PreludeLogger) logWithPrompt(prompt string, s string, args ...any) {
	log.Printf(prompt+" "+s, args...)
}

func (logging *PreludeLogger) Info(s string, args ...any) {
	logging.logWithPrompt(logging.prompt.Info, s, args...)
}

func (logging *PreludeLogger) Warn(s string, args ...any) {
	logging.logWithPrompt(logging.prompt.Warn, s, args...)
}

func (logging *PreludeLogger) Error(s string, args ...any) {
	logging.logWithPrompt(logging.prompt.Error, s, args...)
}

func (logging *PreludeLogger) Debug(s string, args ...any) {
	logging.logWithPrompt(logging.prompt.Debug, s, args...)
}

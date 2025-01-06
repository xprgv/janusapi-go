package logger

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func Null() Logger { return nullLogger{} }

type nullLogger struct{}

func (l nullLogger) Debug(msg string, args ...any) {}
func (l nullLogger) Info(msg string, args ...any)  {}
func (l nullLogger) Warn(msg string, args ...any)  {}
func (l nullLogger) Error(msg string, args ...any) {}

package interfaces

type Logging interface {
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Debug(string, ...any)
}

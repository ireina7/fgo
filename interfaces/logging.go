package interfaces

type Logging interface {
	Log(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Debug(string, ...any)
}

package exception

type Exception interface {
	error
	StackTrace() []*StackInfo
}

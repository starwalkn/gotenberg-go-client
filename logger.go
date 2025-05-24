package gotenberg

type Logger interface {
	Printf(format string, v ...interface{})
}

// NopLogger is an implementation of Logger that does not output messages.
type NopLogger struct {}

func (l NopLogger) Printf(format string, v ...interface{}) {}
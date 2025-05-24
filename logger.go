package gotenberg

type Logger interface {
	Printf(format string, v ...interface{})
}

type NopLogger struct {}

func (l NopLogger) Printf(format string, v ...interface{}) {}
package interfaces

type Logger interface {
	Info(string, ...string)
	Error(string, ...string)
	Warn(string, ...string)
}

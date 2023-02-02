package interfaces

type Logger interface {
	Info(string)
	Infof(string, string)
	Error(string)
	Errorf(string, string)
	Warn(string)
	Warnf(string, string)
}

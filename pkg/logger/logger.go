package logger

type Field map[string]any

type LoggerInterface interface {
	Info(err error)
	Error(err error)
	Debug(err error)
	WithFields(fields Field) LoggerInterface
}

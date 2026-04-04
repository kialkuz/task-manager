package logger

type HttpErrorWriter struct {
	Logger LoggerInterface
}

func (h *HttpErrorWriter) Write(p []byte) (n int, err error) {
	h.Logger.Error(err)
	return len(p), nil
}

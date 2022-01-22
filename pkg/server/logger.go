package server

import (
	"github.com/sirupsen/logrus"
	"io"
)

type DefaultLogger struct {
	w io.Writer
}

func NewLogger(w io.Writer) *DefaultLogger {
	return &DefaultLogger{
		w: w,
	}
}

func (d *DefaultLogger) Printf(format string, args ...interface{})  {
	logrus.SetOutput(d.w)
	logrus.Printf(format, args)
}
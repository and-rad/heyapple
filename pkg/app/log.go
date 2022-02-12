package app

import (
	"fmt"
	"io"
)

type Logger interface {
	Log(interface{})
	Warn(interface{})
	Error(interface{})
}

type Log struct {
	out io.Writer
}

func NewLog(wr io.Writer) *Log {
	return &Log{
		out: wr,
	}
}

func (l *Log) Log(i interface{}) {
	l.send(i, "Log")
}

func (l *Log) Warn(i interface{}) {
	l.send(i, "Warn")
}

func (l *Log) Error(i interface{}) {
	l.send(i, "Error")
}

func (l *Log) send(data interface{}, level string) {
	switch v := data.(type) {
	case string:
		fmt.Fprintf(l.out, "%-5s: %s\n", level, v)
	case error:
		fmt.Fprintf(l.out, "%-5s: %s\n", level, v.Error())
	default:
		fmt.Fprintf(l.out, "%-5s: %v\n", level, v)
	}
}

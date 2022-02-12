package mock

type Log struct {
	Message string
	Warning string
	Err     string
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Log(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Message = v
	case error:
		l.Message = v.Error()
	}
}

func (l *Log) Warn(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Warning = v
	case error:
		l.Warning = v.Error()
	}
}

func (l *Log) Error(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Err = v
	case error:
		l.Err = v.Error()
	}
}

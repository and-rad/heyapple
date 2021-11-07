package mock

import (
	"errors"
)

// Error definitions
var (
	ErrMarshalFail = errors.New("fail")
)

type MarshalFailer struct{}

func (t MarshalFailer) MarshalJSON() ([]byte, error) {
	return nil, ErrMarshalFail
}

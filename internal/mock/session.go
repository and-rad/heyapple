package mock

import (
	"context"
	"errors"

	"github.com/and-rad/scs/v2/memstore"
)

var (
	ErrFailSessionDestroy = errors.New("faildestroy")
)

type SessionStore struct {
	memstore.MemStore
	failDestroy bool
}

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (s *SessionStore) WithFailDestroy() *SessionStore {
	s.failDestroy = true
	return s
}

func (s *SessionStore) DeleteCtx(ctx context.Context, token string) error {
	if s.failDestroy {
		return ErrFailSessionDestroy
	}
	return s.Delete(token)
}

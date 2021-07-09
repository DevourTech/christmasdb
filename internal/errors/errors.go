package errors

import "github.com/pkg/errors"

var (
	// ErrKeyNotFound is the error returned when a key is not found in the store
	ErrKeyNotFound = errors.New("key not found in the store")
)

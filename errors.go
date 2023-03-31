package godel

import "errors"

var (
	ErrEventNotDefined = errors.New("event not defined")
	ErrStateNotDefined = errors.New("state not defined")
)

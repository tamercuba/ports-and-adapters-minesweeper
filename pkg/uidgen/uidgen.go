package uidgen

import "github.com/google/uuid"

// UIDGen interface
type UIDGen interface {
	New() string
}

type uidgen struct{}

// New is the UIDGen factory
func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() string {
	return uuid.New().String()
}

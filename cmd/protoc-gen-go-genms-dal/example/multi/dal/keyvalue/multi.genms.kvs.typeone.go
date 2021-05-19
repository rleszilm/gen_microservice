// Package keyvalue_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package keyvalue_dal_multi

import (
	context "context"

	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
)

// TypeOneKey defines a Key in the kv store.
type TypeOneKey interface{}

// TypeOneReader is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeOneReader
type TypeOneReader interface {
	GetByKey(context.Context, TypeOneKey) (*multi.TypeOne, error)
}

// TypeOneReadeAllr is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeOneReadAller
type TypeOneReadAller interface {
	TypeOneReader
	All(context.Context) ([]*multi.TypeOne, error)
}

// TypeOneWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeOneWriter
type TypeOneWriter interface {
	SetByKey(context.Context, TypeOneKey, *multi.TypeOne) error
}

// TypeOneReadWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeOneReadWriter
type TypeOneReadWriter interface {
	TypeOneReader
	TypeOneWriter
}

// TypeOneKeyFunc is a function that generates a unique deterministic key for the multi.TypeOne.
type TypeOneKeyFunc func(*multi.TypeOne) interface{}

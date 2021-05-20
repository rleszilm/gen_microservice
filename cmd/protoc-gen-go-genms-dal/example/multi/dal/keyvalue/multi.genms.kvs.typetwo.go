// Package keyvalue_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package keyvalue_dal_multi

import (
	context "context"

	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
)

// TypeTwoKey defines a Key in the kv store.
type TypeTwoKey interface{}

// TypeTwoReader is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeTwoReader
type TypeTwoReader interface {
	GetByKey(context.Context, TypeTwoKey) (*multi.TypeTwo, error)
}

// TypeTwoReadeAllr is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeTwoReadAller
type TypeTwoReadAller interface {
	All(context.Context) ([]*multi.TypeTwo, error)
}

// TypeTwoWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeTwoWriter
type TypeTwoWriter interface {
	SetByKey(context.Context, TypeTwoKey, *multi.TypeTwo) error
}

// TypeTwoReadWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TypeTwoReadWriter
type TypeTwoReadWriter interface {
	TypeTwoReader
	TypeTwoWriter
}

// TypeTwoKeyFunc is a function that generates a unique deterministic key for the multi.TypeTwo.
type TypeTwoKeyFunc func(*multi.TypeTwo) interface{}

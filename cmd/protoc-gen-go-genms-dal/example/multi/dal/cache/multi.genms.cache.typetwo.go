// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"

	cache "github.com/rleszilm/genms/cache"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
)

// TypeTwoKey defines a Key in the cache.
type TypeTwoKey interface{}

// TypeTwoReader is defines the interface for getting values from a cache.
//counterfeiter:generate . TypeTwoReader
type TypeTwoReader interface {
	Get(ctx context.Context, key TypeTwoKey) (*multi.TypeTwo, error)
}

// TypeTwoReadeAller is defines the interface for getting all values from a cache.
//counterfeiter:generate . TypeTwoReadAller
type TypeTwoReadAller interface {
	All(ctx context.Context) ([]*multi.TypeTwo, error)
}

// TypeTwoWriter is defines the interface for setting values in a cache.
//counterfeiter:generate . TypeTwoWriter
type TypeTwoWriter interface {
	Set(ctx context.Context, key TypeTwoKey, obj *multi.TypeTwo) (*multi.TypeTwo, error)
}

// TypeTwoReadWriter is defines the interface for setting values in a cache.
//counterfeiter:generate . TypeTwoReadWriter
type TypeTwoReadWriter interface {
	TypeTwoReader
	TypeTwoWriter
}

// TypeTwoKeyFunc is a function that generates a unique deterministic key for the multi.TypeTwo.
type TypeTwoKeyFunc func(*multi.TypeTwo) interface{}

// UnimplementedTypeTwoCache is a KV ReadWriter that takes no action on read or write.
type UnimplementedTypeTwoCache struct {
}

// GetAll implements TypeTwoReadAller.
func (x *UnimplementedTypeTwoCache) All(_ context.Context) (*multi.TypeTwo, error) {
	return nil, cache.ErrUnimplemented
}

// GetByKey implements TypeTwoReader.
func (x *UnimplementedTypeTwoCache) GetByKey(_ context.Context, _ TypeTwoKey) (*multi.TypeTwo, error) {
	return nil, cache.ErrUnimplemented
}

// SetByKey implements TypeTwoWriter.
func (x *UnimplementedTypeTwoCache) SetByKey(_ context.Context, _ TypeTwoKey, _ *multi.TypeTwo) (*multi.TypeTwo, error) {
	return nil, cache.ErrUnimplemented
}

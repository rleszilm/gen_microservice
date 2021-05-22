// Package cache_dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_single

import (
	context "context"

	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/keyvalue"
)

// NilSingleCache is a KV ReadWriter that takes no action on read or write.
type NilSingleCache struct {
}

// GetAll implements keyvalue.SingleReadAller.
func (x *NilSingleCache) GetAll(_ context.Context) (*single.Single, error) {
	return nil, nil
}

// GetByKey implements keyvalue.SingleReader.
func (x *NilSingleCache) GetByKey(_ context.Context, _ keyvalue.SingleKey) (*single.Single, error) {
	return nil, nil
}

// SetByKey implements keyvalue.SingleWriter.
func (x *NilSingleCache) SetByKey(_ context.Context, _ keyvalue.SingleKey, _ *single.Single) error {
	return nil
}

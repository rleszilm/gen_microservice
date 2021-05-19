// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"
	time "time"

	cache "github.com/rleszilm/genms/cache"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeOneMap defines a Map base cache implementing keyvalue.TypeOneReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type TypeOneMap struct {
	name   string
	reader keyvalue.TypeOneReader
	writer keyvalue.TypeOneWriter
	cache  map[keyvalue.TypeOneKey]*multi.TypeOne
	all    []*multi.TypeOne
}

// All implements implements keyvalue.TypeOneReadAller.
func (x *TypeOneMap) All(ctx context.Context) ([]*multi.TypeOne, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCacheCollection, "type_one"),
		tag.Upsert(cache.TagCacheInstance, x.name),
		tag.Upsert(cache.TagCacheMethod, "all"),
		tag.Upsert(cache.TagCacheType, "map"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	return x.all, nil
}

// GetByKey implements keyvalue.TypeOneReader.
func (x *TypeOneMap) GetByKey(ctx context.Context, key keyvalue.TypeOneKey) (*multi.TypeOne, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCacheCollection, "type_one"),
		tag.Upsert(cache.TagCacheInstance, x.name),
		tag.Upsert(cache.TagCacheMethod, "get"),
		tag.Upsert(cache.TagCacheType, "map"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	if val, ok := x.cache[key]; ok {
		return val, nil
	}

	if x.reader != nil {
		val, err := x.reader.GetByKey(ctx, key)
		if err != nil {
			x.cache[key] = val
			return val, nil
		}
	}

	stats.Record(ctx, cache.MeasureError.M(1))
	return nil, nil
}

// SetByKey implements keyvalue.TypeOneWriter.
func (x *TypeOneMap) SetByKey(ctx context.Context, key keyvalue.TypeOneKey, val *multi.TypeOne) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCacheCollection, "type_one"),
		tag.Upsert(cache.TagCacheInstance, x.name),
		tag.Upsert(cache.TagCacheMethod, "get"),
		tag.Upsert(cache.TagCacheType, "map"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	if x.writer != nil {
		if err := x.writer.SetByKey(ctx, key, val); err != nil {
			stats.Record(ctx, cache.MeasureError.M(1))
			return err
		}
	}

	x.cache[key] = val
	return nil
}

// WithReader tells the TypeOneMap where to source values from if they don't exist in cache.
func (x *TypeOneMap) WithReader(r keyvalue.TypeOneReader) {
	x.reader = r
}

// WithWriter tells the TypeOneMap where to source values from if they don't exist in cache.
func (x *TypeOneMap) WithWriter(w keyvalue.TypeOneWriter) {
	x.writer = w
}

// NewTypeOneMap returns a new TypeOneMap cache.
func NewTypeOneMap(name string, i int) (*TypeOneMap, error) {
	return &TypeOneMap{
		name:  name,
		cache: map[keyvalue.TypeOneKey]*multi.TypeOne{},
	}, nil
}

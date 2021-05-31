// Package cache_dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_single

import (
	context "context"
	fmt "fmt"
	time "time"

	cache "github.com/rleszilm/genms/cache"
	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/keyvalue"
	service "github.com/rleszilm/genms/service"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// SingleMap defines a Map base cache implementing keyvalue.SingleReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type SingleMap struct {
	service.Dependencies
	NilSingleCache

	name   string
	reader keyvalue.SingleReader
	writer keyvalue.SingleWriter
	cache  map[keyvalue.SingleKey]*single.Single
	all    []*single.Single
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *SingleMap) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *SingleMap) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of the map.
func (x *SingleMap) NameOf() string {
	return x.name
}

// String returns the name of the map.
func (x *SingleMap) String() string {
	return x.name
}

// All implements implements keyvalue.SingleReadAller.
func (x *SingleMap) All(ctx context.Context) ([]*single.Single, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "all"),
		tag.Upsert(cache.TagType, "map"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	return x.all, nil
}

// GetByKey implements keyvalue.SingleReader.
func (x *SingleMap) GetByKey(ctx context.Context, key keyvalue.SingleKey) (*single.Single, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "get"),
		tag.Upsert(cache.TagType, "map"),
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
			return nil, fmt.Errorf("map: Single.GetByKey - %w", err)
		}
		x.cache[key] = val
		return val, nil
	}

	stats.Record(ctx, cache.MeasureError.M(1))
	return nil, fmt.Errorf("map: Single.GetByKey - %w", cache.ErrGetValue)
}

// SetByKey implements keyvalue.SingleWriter.
func (x *SingleMap) SetByKey(ctx context.Context, key keyvalue.SingleKey, val *single.Single) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "get"),
		tag.Upsert(cache.TagType, "map"),
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
			return fmt.Errorf("map: Single.SetByKey - %w", err)
		}
	}

	x.cache[key] = val

	all := []*single.Single{}
	for _, v := range x.cache {
		all = append(all, v)
	}
	x.all = all

	return nil
}

// WithReader tells the SingleMap where to source values from if they don't exist in cache.
func (x *SingleMap) WithReader(r keyvalue.SingleReader) {
	x.reader = r
}

// WithWriter tells the SingleMap where to source values from if they don't exist in cache.
func (x *SingleMap) WithWriter(w keyvalue.SingleWriter) {
	x.writer = w
}

// NewSingleMap returns a new SingleMap cache.
func NewSingleMap(name string) (*SingleMap, error) {
	return &SingleMap{
		name:  name,
		cache: map[keyvalue.SingleKey]*single.Single{},
	}, nil
}

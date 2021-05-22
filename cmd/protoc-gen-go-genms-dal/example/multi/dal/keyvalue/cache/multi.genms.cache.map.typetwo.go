// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"
	fmt "fmt"
	time "time"

	cache "github.com/rleszilm/genms/cache"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
	service "github.com/rleszilm/genms/service"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeTwoMap defines a Map base cache implementing keyvalue.TypeTwoReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type TypeTwoMap struct {
	service.Dependencies
	NilTypeTwoCache

	name   string
	reader keyvalue.TypeTwoReader
	writer keyvalue.TypeTwoWriter
	cache  map[keyvalue.TypeTwoKey]*multi.TypeTwo
	all    []*multi.TypeTwo
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *TypeTwoMap) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *TypeTwoMap) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of the map.
func (x *TypeTwoMap) NameOf() string {
	return x.name
}

// String returns the name of the map.
func (x *TypeTwoMap) String() string {
	return x.name
}

// All implements implements keyvalue.TypeTwoReadAller.
func (x *TypeTwoMap) All(ctx context.Context) ([]*multi.TypeTwo, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
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

// GetByKey implements keyvalue.TypeTwoReader.
func (x *TypeTwoMap) GetByKey(ctx context.Context, key keyvalue.TypeTwoKey) (*multi.TypeTwo, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
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
			return nil, fmt.Errorf("map: <no value>.GetByKey - %w", err)
		}
		x.cache[key] = val
		return val, nil
	}

	stats.Record(ctx, cache.MeasureError.M(1))
	return nil, fmt.Errorf("map: <no value>.GetByKey - %w", cache.ErrGetValue)
}

// SetByKey implements keyvalue.TypeTwoWriter.
func (x *TypeTwoMap) SetByKey(ctx context.Context, key keyvalue.TypeTwoKey, val *multi.TypeTwo) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
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
			return fmt.Errorf("map: <no value>.SetByKey - %w", err)
		}
	}

	x.cache[key] = val
	return nil
}

// WithReader tells the TypeTwoMap where to source values from if they don't exist in cache.
func (x *TypeTwoMap) WithReader(r keyvalue.TypeTwoReader) {
	x.reader = r
}

// WithWriter tells the TypeTwoMap where to source values from if they don't exist in cache.
func (x *TypeTwoMap) WithWriter(w keyvalue.TypeTwoWriter) {
	x.writer = w
}

// NewTypeTwoMap returns a new TypeTwoMap cache.
func NewTypeTwoMap(name string) (*TypeTwoMap, error) {
	return &TypeTwoMap{
		name:  name,
		cache: map[keyvalue.TypeTwoKey]*multi.TypeTwo{},
	}, nil
}

// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"
	time "time"

	golang_lru "github.com/hashicorp/golang-lru"
	cache "github.com/rleszilm/genms/cache"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeTwoLRU defines a LRU cache implementing keyvalue.TypeTwoReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type TypeTwoLRU struct {
	name   string
	reader keyvalue.TypeTwoReader
	writer keyvalue.TypeTwoWriter
	lru    *golang_lru.ARCCache
	all    []*multi.TypeTwo
}

// All implements implements keyvalue.TypeTwoReadAller.
func (x *TypeTwoLRU) All(ctx context.Context) ([]*multi.TypeTwo, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "all"),
		tag.Upsert(cache.TagType, "lru"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	return x.all, nil
}

// GetByKey implements keyvalue.TypeTwoReadWriter.
func (x *TypeTwoLRU) GetByKey(ctx context.Context, key keyvalue.TypeTwoKey) (*multi.TypeTwo, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "get"),
		tag.Upsert(cache.TagType, "lru"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	if val, ok := x.lru.Get(key); ok {
		stats.Record(ctx, cache.MeasureHit.M(1))
		return val.(*multi.TypeTwo), nil
	}
	stats.Record(ctx, cache.MeasureMiss.M(1))

	if x.reader != nil {
		val, err := x.reader.GetByKey(ctx, key)
		if err != nil {
			return nil, err
		}
		x.lru.Add(key, val)
		return val, nil
	}

	stats.Record(ctx, cache.MeasureError.M(1))
	return nil, nil
}

// SetByKey implements keyvalue.TypeTwoReadWriter.
func (x *TypeTwoLRU) SetByKey(ctx context.Context, key keyvalue.TypeTwoKey, val *multi.TypeTwo) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "set"),
		tag.Upsert(cache.TagType, "lru"),
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

	x.lru.Add(key, val)

	all := []*multi.TypeTwo{}
	for _, k := range x.lru.Keys() {
		y, _ := x.lru.Get(k)
		all = append(all, y.(*multi.TypeTwo))
	}
	x.all = all

	return nil
}

// WithReader tells the TypeTwoLRU where to source values from if they don't exist in cache.
func (x *TypeTwoLRU) WithReader(r keyvalue.TypeTwoReader) {
	x.reader = r
}

// WithWriter tells the TypeTwoLRU where to source values from if they don't exist in cache.
func (x *TypeTwoLRU) WithWriter(w keyvalue.TypeTwoWriter) {
	x.writer = w
}

// NewTypeTwoLRU returns a new TypeTwoLRU cache.
func NewTypeTwoLRU(name string, i int) (*TypeTwoLRU, error) {
	arc, err := golang_lru.NewARC(i)
	if err != nil {
		return nil, err
	}

	return &TypeTwoLRU{
		name: name,
		lru:  arc,
	}, nil
}

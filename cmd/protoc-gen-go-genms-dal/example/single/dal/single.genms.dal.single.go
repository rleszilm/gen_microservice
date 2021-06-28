// Package dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package dal_single

import (
	context "context"
	errors "errors"
	fmt "fmt"
	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	service "github.com/rleszilm/genms/service"
)

var (
	// ErrSingleCollectionMethodImpl is returned when the called method is not implemented.
	ErrSingleCollectionMethodImpl = errors.New("SingleCollection method is not implemented")
)

// SingleMutator is a function that makes changes to a Single.
type SingleMutator func(*single.Single) (*single.Single, error)

// SingleCollection is an autogenerated interface that can be used to interact with a collection of Single objects.
//counterfeiter:generate . SingleCollection
type SingleCollection interface {
	service.Service

	SingleCollectionReader
	SingleCollectionWriter
}

// SingleCollectionWriter is an autogenerated interface that can be used to write to a collection of Single objects.
//counterfeiter:generate . SingleCollectionWriter
type SingleCollectionWriter interface {
	// Insert runs the command to generate a new object within the data store.
	Insert(ctx context.Context, obj *single.Single) (*single.Single, error)
	// Upsert runs the command to overwrite the object in the datastore, or write it if it does not already exist.
	Upsert(ctx context.Context, obj *single.Single) (*single.Single, error)
	// Update runs the command to make changes to the given record.
	Update(ctx context.Context, obj *single.Single, update *SingleFieldValues) (*single.Single, error)
}

// SingleCollectionReader is an autogenerated interface that can be used to query a collection
// of Single objects. The queries and their values are taken from the representative proto message.
//counterfeiter:generate . SingleCollectionReader
type SingleCollectionReader interface {
	All(ctx context.Context) ([]*single.Single, error)
	Filter(ctx context.Context, filter *SingleFieldValues) ([]*single.Single, error)
	ById(ctx context.Context, bsonStringOid string) ([]*single.Single, error)
	OneParam(ctx context.Context, scalarInt32 int32) ([]*single.Single, error)
	MultipleParam(ctx context.Context, scalarInt32 int32, scalarInt64 int64, scalarFloat32 float32) ([]*single.Single, error)
	MessageParam(ctx context.Context, objMessage *single.Single_Message) ([]*single.Single, error)
	WithComparator(ctx context.Context, scalarInt32 int32) ([]*single.Single, error)
	WithRest(ctx context.Context, scalarInt32 int32, scalarInt64 int64, scalarFloat32 float32, scalarFloat64 float64) ([]*single.Single, error)
	ProviderStubOnly(ctx context.Context) ([]*single.Single, error)
	InterfaceStubOnly(ctx context.Context) ([]*single.Single, error)
	NonFieldOnly(ctx context.Context, kind string) ([]*single.Single, error)
}

// SingleFieldValues is an autogenerated struct that can be used in the generic queries against SingleCollection.
type SingleFieldValues struct {
	ScalarInt32     *int32
	ScalarInt64     *int64
	ScalarFloat32   *float32
	ScalarFloat64   *float64
	ScalarString    *string
	ScalarBytes     []byte
	ScalarBool      *bool
	ManyScalarBool  []bool
	ScalarEnum      *single.Single_Enum
	ObjMessage      *single.Single_Message
	ManyObjMessage  []*single.Single_Message
	Renamed         *string
	IgnoredPostgres *string
	RenamedPostgres *string
	IgnoredRest     *string
	RenamedRest     *string
	IgnoredMongo    *string
	RenamedMongo    *string
	BsonStringOid   *string
	BsonBytesOid    []byte
}

// UnimplementedSingleCollection is an autogenerated implementation of SingleCollection that returns an error when any
// method is called.
type UnimplementedSingleCollection struct {
	service.Dependencies
}

// Insert implements SingleCollection.Insert
func (x *UnimplementedSingleCollection) Insert(_ context.Context, _ *single.Single) (*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// Upsert implements SingleCollection.Upsert
func (x *UnimplementedSingleCollection) Upsert(_ context.Context, _ *single.Single) (*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// Update implements SingleCollection.Update
func (x *UnimplementedSingleCollection) Update(_ context.Context, _ *single.Single, _ *SingleFieldValues) (*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// Filter implements SingleCollection.Filter
func (x *UnimplementedSingleCollection) Filter(_ context.Context, _ *SingleFieldValues) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// ById implements SingleCollection.ById
func (x *UnimplementedSingleCollection) ById(_ context.Context, _ string) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// OneParam implements SingleCollection.OneParam
func (x *UnimplementedSingleCollection) OneParam(_ context.Context, _ int32) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// MultipleParam implements SingleCollection.MultipleParam
func (x *UnimplementedSingleCollection) MultipleParam(_ context.Context, _ int32, _ int64, _ float32) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// MessageParam implements SingleCollection.MessageParam
func (x *UnimplementedSingleCollection) MessageParam(_ context.Context, _ *single.Single_Message) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// WithComparator implements SingleCollection.WithComparator
func (x *UnimplementedSingleCollection) WithComparator(_ context.Context, _ int32) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// WithRest implements SingleCollection.WithRest
func (x *UnimplementedSingleCollection) WithRest(_ context.Context, _ int32, _ int64, _ float32, _ float64) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// ProviderStubOnly implements SingleCollection.ProviderStubOnly
func (x *UnimplementedSingleCollection) ProviderStubOnly(_ context.Context) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// InterfaceStubOnly implements SingleCollection.InterfaceStubOnly
func (x *UnimplementedSingleCollection) InterfaceStubOnly(_ context.Context) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

// NonFieldOnly implements SingleCollection.NonFieldOnly
func (x *UnimplementedSingleCollection) NonFieldOnly(_ context.Context, _ string) ([]*single.Single, error) {
	return nil, ErrSingleCollectionMethodImpl
}

func ReturnsOneSingle(xs []*single.Single, err error) (*single.Single, error) {
	if err != nil {
		return nil, err
	}

	switch len(xs) {
	case 0:
		return nil, err
	case 1:
		return xs[0], err
	default:
		return nil, fmt.Errorf("single.single: more than 1 value returned - %w", err)
	}
}

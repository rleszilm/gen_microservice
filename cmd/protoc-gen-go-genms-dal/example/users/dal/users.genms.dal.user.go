// Package dal_users is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package dal_users

import (
	context "context"
	errors "errors"
	types "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations/types"
	users "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/users"
	service "github.com/rleszilm/genms/service"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
)

var (
	// ErrUserCollectionMethodImpl is returned when the called method is not implemented.
	ErrUserCollectionMethodImpl = errors.New("UserCollection method is not implemented")
)

// UserCollection is an autogenerated interface that can be used to interact with a collection of User objects.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UserCollection
type UserCollection interface {
	service.Service

	UserCollectionReader
	UserCollectionWriter
}

// UserCollectionWriter is an autogenerated interface that can be used to write to a collection of User objects.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UserCollectionWriter
type UserCollectionWriter interface {
	// Insert runs the command to generate a new object within the data store.
	Insert(context.Context, *users.User) (*users.User, error)
	// Upsert runs the command to overwrite the object in the datastore, or write it if it does nto already exist.
	Upsert(context.Context, *users.User) (*users.User, error)
	// Update runs the command to make changes to the given record.
	Update(context.Context, *UserFieldValues) (*users.User, error)
}

// UserCollectionReader is an autogenerated interface that can be used to query a collection
// of User objects. The queries and their values are taken from the representative proto message.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UserCollectionReader
type UserCollectionReader interface {
	All(context.Context) ([]*users.User, error)
	Filter(context.Context, *UserFieldValues) ([]*users.User, error)
	ById(_ context.Context, _ int64) ([]*users.User, error)
	ByNameAndDivision(_ context.Context, _ string, _ string) ([]*users.User, error)
	ByKind(_ context.Context, _ users.User_Kind) ([]*users.User, error)
	ByPhone(_ context.Context, _ *types.Phone) ([]*users.User, error)
	ProviderStubOnly(_ context.Context) ([]*users.User, error)
	InterfaceStubOnly(_ context.Context) ([]*users.User, error)
}

// UserFieldValues is an autogenerated struct that can be used in the generic queries against UserCollection.
type UserFieldValues struct {
	Id            *int64
	Name          *string
	Division      *string
	LifetimeScore *float64
	LastScore     *float32
	LastWinnings  *int32
	Point         *types.Point
	Phone         *types.Phone
	Geo           *latlng.LatLng
	Kind          *users.User_Kind
	ByBackend     *string
}

// UnimplementedUserCollection is an autogenerated implementation of UserCollection that returns an error when any
// method is called.
type UnimplementedUserCollection struct {
	service.Dependencies
}

// Insert implements UserCollection.Insert
func (x *UnimplementedUserCollection) Insert(_ context.Context, _ *users.User) (*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// Upsert implements UserCollection.Upsert
func (x *UnimplementedUserCollection) Upsert(_ context.Context, _ *users.User) (*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// Update implements UserCollection.Update
func (x *UnimplementedUserCollection) Update(_ context.Context, _ *UserFieldValues) (*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// Filter implements UserCollection.Filter
func (x *UnimplementedUserCollection) Filter(_ context.Context, _ *UserFieldValues) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ById implements UserCollection.ById
func (x *UnimplementedUserCollection) ById(_ context.Context, _ int64) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ByNameAndDivision implements UserCollection.ByNameAndDivision
func (x *UnimplementedUserCollection) ByNameAndDivision(_ context.Context, _ string, _ string) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ByKind implements UserCollection.ByKind
func (x *UnimplementedUserCollection) ByKind(_ context.Context, _ users.User_Kind) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ByPhone implements UserCollection.ByPhone
func (x *UnimplementedUserCollection) ByPhone(_ context.Context, _ *types.Phone) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ProviderStubOnly implements UserCollection.ProviderStubOnly
func (x *UnimplementedUserCollection) ProviderStubOnly(_ context.Context) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// InterfaceStubOnly implements UserCollection.InterfaceStubOnly
func (x *UnimplementedUserCollection) InterfaceStubOnly(_ context.Context) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

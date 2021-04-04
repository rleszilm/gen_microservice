// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"context"
	"sync"

	usersa "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/users"
	users "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/users/dal"
)

type FakeUserCollectionWriter struct {
	InsertStub        func(context.Context, *usersa.User) (*usersa.User, error)
	insertMutex       sync.RWMutex
	insertArgsForCall []struct {
		arg1 context.Context
		arg2 *usersa.User
	}
	insertReturns struct {
		result1 *usersa.User
		result2 error
	}
	insertReturnsOnCall map[int]struct {
		result1 *usersa.User
		result2 error
	}
	UpsertStub        func(context.Context, *usersa.User) (*usersa.User, error)
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		arg1 context.Context
		arg2 *usersa.User
	}
	upsertReturns struct {
		result1 *usersa.User
		result2 error
	}
	upsertReturnsOnCall map[int]struct {
		result1 *usersa.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserCollectionWriter) Insert(arg1 context.Context, arg2 *usersa.User) (*usersa.User, error) {
	fake.insertMutex.Lock()
	ret, specificReturn := fake.insertReturnsOnCall[len(fake.insertArgsForCall)]
	fake.insertArgsForCall = append(fake.insertArgsForCall, struct {
		arg1 context.Context
		arg2 *usersa.User
	}{arg1, arg2})
	fake.recordInvocation("Insert", []interface{}{arg1, arg2})
	fake.insertMutex.Unlock()
	if fake.InsertStub != nil {
		return fake.InsertStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.insertReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionWriter) InsertCallCount() int {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return len(fake.insertArgsForCall)
}

func (fake *FakeUserCollectionWriter) InsertCalls(stub func(context.Context, *usersa.User) (*usersa.User, error)) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = stub
}

func (fake *FakeUserCollectionWriter) InsertArgsForCall(i int) (context.Context, *usersa.User) {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	argsForCall := fake.insertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionWriter) InsertReturns(result1 *usersa.User, result2 error) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = nil
	fake.insertReturns = struct {
		result1 *usersa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionWriter) InsertReturnsOnCall(i int, result1 *usersa.User, result2 error) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = nil
	if fake.insertReturnsOnCall == nil {
		fake.insertReturnsOnCall = make(map[int]struct {
			result1 *usersa.User
			result2 error
		})
	}
	fake.insertReturnsOnCall[i] = struct {
		result1 *usersa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionWriter) Upsert(arg1 context.Context, arg2 *usersa.User) (*usersa.User, error) {
	fake.upsertMutex.Lock()
	ret, specificReturn := fake.upsertReturnsOnCall[len(fake.upsertArgsForCall)]
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		arg1 context.Context
		arg2 *usersa.User
	}{arg1, arg2})
	fake.recordInvocation("Upsert", []interface{}{arg1, arg2})
	fake.upsertMutex.Unlock()
	if fake.UpsertStub != nil {
		return fake.UpsertStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.upsertReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionWriter) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *FakeUserCollectionWriter) UpsertCalls(stub func(context.Context, *usersa.User) (*usersa.User, error)) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = stub
}

func (fake *FakeUserCollectionWriter) UpsertArgsForCall(i int) (context.Context, *usersa.User) {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	argsForCall := fake.upsertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionWriter) UpsertReturns(result1 *usersa.User, result2 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	fake.upsertReturns = struct {
		result1 *usersa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionWriter) UpsertReturnsOnCall(i int, result1 *usersa.User, result2 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	if fake.upsertReturnsOnCall == nil {
		fake.upsertReturnsOnCall = make(map[int]struct {
			result1 *usersa.User
			result2 error
		})
	}
	fake.upsertReturnsOnCall[i] = struct {
		result1 *usersa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserCollectionWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ users.UserCollectionWriter = new(FakeUserCollectionWriter)

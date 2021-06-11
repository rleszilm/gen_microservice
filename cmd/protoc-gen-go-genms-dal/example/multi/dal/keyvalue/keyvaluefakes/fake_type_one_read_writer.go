// Code generated by counterfeiter. DO NOT EDIT.
package keyvaluefakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue_dal_multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
)

type FakeTypeOneReadWriter struct {
	GetByKeyStub        func(context.Context, keyvalue_dal_multi.TypeOneKey) (*multi.TypeOne, error)
	getByKeyMutex       sync.RWMutex
	getByKeyArgsForCall []struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeOneKey
	}
	getByKeyReturns struct {
		result1 *multi.TypeOne
		result2 error
	}
	getByKeyReturnsOnCall map[int]struct {
		result1 *multi.TypeOne
		result2 error
	}
	SetByKeyStub        func(context.Context, keyvalue_dal_multi.TypeOneKey, *multi.TypeOne) error
	setByKeyMutex       sync.RWMutex
	setByKeyArgsForCall []struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeOneKey
		arg3 *multi.TypeOne
	}
	setByKeyReturns struct {
		result1 error
	}
	setByKeyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypeOneReadWriter) GetByKey(arg1 context.Context, arg2 keyvalue_dal_multi.TypeOneKey) (*multi.TypeOne, error) {
	fake.getByKeyMutex.Lock()
	ret, specificReturn := fake.getByKeyReturnsOnCall[len(fake.getByKeyArgsForCall)]
	fake.getByKeyArgsForCall = append(fake.getByKeyArgsForCall, struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeOneKey
	}{arg1, arg2})
	fake.recordInvocation("GetByKey", []interface{}{arg1, arg2})
	fake.getByKeyMutex.Unlock()
	if fake.GetByKeyStub != nil {
		return fake.GetByKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getByKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeOneReadWriter) GetByKeyCallCount() int {
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	return len(fake.getByKeyArgsForCall)
}

func (fake *FakeTypeOneReadWriter) GetByKeyCalls(stub func(context.Context, keyvalue_dal_multi.TypeOneKey) (*multi.TypeOne, error)) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = stub
}

func (fake *FakeTypeOneReadWriter) GetByKeyArgsForCall(i int) (context.Context, keyvalue_dal_multi.TypeOneKey) {
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	argsForCall := fake.getByKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeOneReadWriter) GetByKeyReturns(result1 *multi.TypeOne, result2 error) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = nil
	fake.getByKeyReturns = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneReadWriter) GetByKeyReturnsOnCall(i int, result1 *multi.TypeOne, result2 error) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = nil
	if fake.getByKeyReturnsOnCall == nil {
		fake.getByKeyReturnsOnCall = make(map[int]struct {
			result1 *multi.TypeOne
			result2 error
		})
	}
	fake.getByKeyReturnsOnCall[i] = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneReadWriter) SetByKey(arg1 context.Context, arg2 keyvalue_dal_multi.TypeOneKey, arg3 *multi.TypeOne) error {
	fake.setByKeyMutex.Lock()
	ret, specificReturn := fake.setByKeyReturnsOnCall[len(fake.setByKeyArgsForCall)]
	fake.setByKeyArgsForCall = append(fake.setByKeyArgsForCall, struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeOneKey
		arg3 *multi.TypeOne
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetByKey", []interface{}{arg1, arg2, arg3})
	fake.setByKeyMutex.Unlock()
	if fake.SetByKeyStub != nil {
		return fake.SetByKeyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setByKeyReturns
	return fakeReturns.result1
}

func (fake *FakeTypeOneReadWriter) SetByKeyCallCount() int {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	return len(fake.setByKeyArgsForCall)
}

func (fake *FakeTypeOneReadWriter) SetByKeyCalls(stub func(context.Context, keyvalue_dal_multi.TypeOneKey, *multi.TypeOne) error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = stub
}

func (fake *FakeTypeOneReadWriter) SetByKeyArgsForCall(i int) (context.Context, keyvalue_dal_multi.TypeOneKey, *multi.TypeOne) {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	argsForCall := fake.setByKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTypeOneReadWriter) SetByKeyReturns(result1 error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = nil
	fake.setByKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTypeOneReadWriter) SetByKeyReturnsOnCall(i int, result1 error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = nil
	if fake.setByKeyReturnsOnCall == nil {
		fake.setByKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setByKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTypeOneReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypeOneReadWriter) recordInvocation(key string, args []interface{}) {
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

var _ keyvalue_dal_multi.TypeOneReadWriter = new(FakeTypeOneReadWriter)
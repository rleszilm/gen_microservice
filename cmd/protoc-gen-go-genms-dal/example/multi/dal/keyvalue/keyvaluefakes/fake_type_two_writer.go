// Code generated by counterfeiter. DO NOT EDIT.
package keyvaluefakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue_dal_multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
)

type FakeTypeTwoWriter struct {
	SetByKeyStub        func(context.Context, keyvalue_dal_multi.TypeTwoKey, *multi.TypeTwo) error
	setByKeyMutex       sync.RWMutex
	setByKeyArgsForCall []struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeTwoKey
		arg3 *multi.TypeTwo
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

func (fake *FakeTypeTwoWriter) SetByKey(arg1 context.Context, arg2 keyvalue_dal_multi.TypeTwoKey, arg3 *multi.TypeTwo) error {
	fake.setByKeyMutex.Lock()
	ret, specificReturn := fake.setByKeyReturnsOnCall[len(fake.setByKeyArgsForCall)]
	fake.setByKeyArgsForCall = append(fake.setByKeyArgsForCall, struct {
		arg1 context.Context
		arg2 keyvalue_dal_multi.TypeTwoKey
		arg3 *multi.TypeTwo
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

func (fake *FakeTypeTwoWriter) SetByKeyCallCount() int {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	return len(fake.setByKeyArgsForCall)
}

func (fake *FakeTypeTwoWriter) SetByKeyCalls(stub func(context.Context, keyvalue_dal_multi.TypeTwoKey, *multi.TypeTwo) error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = stub
}

func (fake *FakeTypeTwoWriter) SetByKeyArgsForCall(i int) (context.Context, keyvalue_dal_multi.TypeTwoKey, *multi.TypeTwo) {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	argsForCall := fake.setByKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTypeTwoWriter) SetByKeyReturns(result1 error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = nil
	fake.setByKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTypeTwoWriter) SetByKeyReturnsOnCall(i int, result1 error) {
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

func (fake *FakeTypeTwoWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypeTwoWriter) recordInvocation(key string, args []interface{}) {
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

var _ keyvalue_dal_multi.TypeTwoWriter = new(FakeTypeTwoWriter)

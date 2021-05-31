// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	dal_multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal"
)

type FakeTypeTwoCollectionReader struct {
	AllStub        func(context.Context) ([]*multi.TypeTwo, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
		arg1 context.Context
	}
	allReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	FilterStub        func(context.Context, *dal_multi.TypeTwoFieldValues) ([]*multi.TypeTwo, error)
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 context.Context
		arg2 *dal_multi.TypeTwoFieldValues
	}
	filterReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	filterReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	InterfaceStubOnlyStub        func(context.Context) ([]*multi.TypeTwo, error)
	interfaceStubOnlyMutex       sync.RWMutex
	interfaceStubOnlyArgsForCall []struct {
		arg1 context.Context
	}
	interfaceStubOnlyReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	interfaceStubOnlyReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	MessageParamStub        func(context.Context, *multi.TypeTwo_Message) ([]*multi.TypeTwo, error)
	messageParamMutex       sync.RWMutex
	messageParamArgsForCall []struct {
		arg1 context.Context
		arg2 *multi.TypeTwo_Message
	}
	messageParamReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	messageParamReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	MultipleParamStub        func(context.Context, int32, int64, float32) ([]*multi.TypeTwo, error)
	multipleParamMutex       sync.RWMutex
	multipleParamArgsForCall []struct {
		arg1 context.Context
		arg2 int32
		arg3 int64
		arg4 float32
	}
	multipleParamReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	multipleParamReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	OneParamStub        func(context.Context, int32) ([]*multi.TypeTwo, error)
	oneParamMutex       sync.RWMutex
	oneParamArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	oneParamReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	oneParamReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	ProviderStubOnlyStub        func(context.Context) ([]*multi.TypeTwo, error)
	providerStubOnlyMutex       sync.RWMutex
	providerStubOnlyArgsForCall []struct {
		arg1 context.Context
	}
	providerStubOnlyReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	providerStubOnlyReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	WithComparatorStub        func(context.Context, int32) ([]*multi.TypeTwo, error)
	withComparatorMutex       sync.RWMutex
	withComparatorArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	withComparatorReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	withComparatorReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	WithRestStub        func(context.Context, int32, int64, float32, float64) ([]*multi.TypeTwo, error)
	withRestMutex       sync.RWMutex
	withRestArgsForCall []struct {
		arg1 context.Context
		arg2 int32
		arg3 int64
		arg4 float32
		arg5 float64
	}
	withRestReturns struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	withRestReturnsOnCall map[int]struct {
		result1 []*multi.TypeTwo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypeTwoCollectionReader) All(arg1 context.Context) ([]*multi.TypeTwo, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("All", []interface{}{arg1})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.allReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) AllCalls(stub func(context.Context) ([]*multi.TypeTwo, error)) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = stub
}

func (fake *FakeTypeTwoCollectionReader) AllArgsForCall(i int) context.Context {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	argsForCall := fake.allArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTypeTwoCollectionReader) AllReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) AllReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) Filter(arg1 context.Context, arg2 *dal_multi.TypeTwoFieldValues) ([]*multi.TypeTwo, error) {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 context.Context
		arg2 *dal_multi.TypeTwoFieldValues
	}{arg1, arg2})
	fake.recordInvocation("Filter", []interface{}{arg1, arg2})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.filterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) FilterCalls(stub func(context.Context, *dal_multi.TypeTwoFieldValues) ([]*multi.TypeTwo, error)) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = stub
}

func (fake *FakeTypeTwoCollectionReader) FilterArgsForCall(i int) (context.Context, *dal_multi.TypeTwoFieldValues) {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	argsForCall := fake.filterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeTwoCollectionReader) FilterReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) FilterReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnly(arg1 context.Context) ([]*multi.TypeTwo, error) {
	fake.interfaceStubOnlyMutex.Lock()
	ret, specificReturn := fake.interfaceStubOnlyReturnsOnCall[len(fake.interfaceStubOnlyArgsForCall)]
	fake.interfaceStubOnlyArgsForCall = append(fake.interfaceStubOnlyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("InterfaceStubOnly", []interface{}{arg1})
	fake.interfaceStubOnlyMutex.Unlock()
	if fake.InterfaceStubOnlyStub != nil {
		return fake.InterfaceStubOnlyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interfaceStubOnlyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnlyCallCount() int {
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	return len(fake.interfaceStubOnlyArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnlyCalls(stub func(context.Context) ([]*multi.TypeTwo, error)) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = stub
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnlyArgsForCall(i int) context.Context {
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	argsForCall := fake.interfaceStubOnlyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnlyReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = nil
	fake.interfaceStubOnlyReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) InterfaceStubOnlyReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = nil
	if fake.interfaceStubOnlyReturnsOnCall == nil {
		fake.interfaceStubOnlyReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.interfaceStubOnlyReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) MessageParam(arg1 context.Context, arg2 *multi.TypeTwo_Message) ([]*multi.TypeTwo, error) {
	fake.messageParamMutex.Lock()
	ret, specificReturn := fake.messageParamReturnsOnCall[len(fake.messageParamArgsForCall)]
	fake.messageParamArgsForCall = append(fake.messageParamArgsForCall, struct {
		arg1 context.Context
		arg2 *multi.TypeTwo_Message
	}{arg1, arg2})
	fake.recordInvocation("MessageParam", []interface{}{arg1, arg2})
	fake.messageParamMutex.Unlock()
	if fake.MessageParamStub != nil {
		return fake.MessageParamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.messageParamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) MessageParamCallCount() int {
	fake.messageParamMutex.RLock()
	defer fake.messageParamMutex.RUnlock()
	return len(fake.messageParamArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) MessageParamCalls(stub func(context.Context, *multi.TypeTwo_Message) ([]*multi.TypeTwo, error)) {
	fake.messageParamMutex.Lock()
	defer fake.messageParamMutex.Unlock()
	fake.MessageParamStub = stub
}

func (fake *FakeTypeTwoCollectionReader) MessageParamArgsForCall(i int) (context.Context, *multi.TypeTwo_Message) {
	fake.messageParamMutex.RLock()
	defer fake.messageParamMutex.RUnlock()
	argsForCall := fake.messageParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeTwoCollectionReader) MessageParamReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.messageParamMutex.Lock()
	defer fake.messageParamMutex.Unlock()
	fake.MessageParamStub = nil
	fake.messageParamReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) MessageParamReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.messageParamMutex.Lock()
	defer fake.messageParamMutex.Unlock()
	fake.MessageParamStub = nil
	if fake.messageParamReturnsOnCall == nil {
		fake.messageParamReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.messageParamReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) MultipleParam(arg1 context.Context, arg2 int32, arg3 int64, arg4 float32) ([]*multi.TypeTwo, error) {
	fake.multipleParamMutex.Lock()
	ret, specificReturn := fake.multipleParamReturnsOnCall[len(fake.multipleParamArgsForCall)]
	fake.multipleParamArgsForCall = append(fake.multipleParamArgsForCall, struct {
		arg1 context.Context
		arg2 int32
		arg3 int64
		arg4 float32
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("MultipleParam", []interface{}{arg1, arg2, arg3, arg4})
	fake.multipleParamMutex.Unlock()
	if fake.MultipleParamStub != nil {
		return fake.MultipleParamStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.multipleParamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) MultipleParamCallCount() int {
	fake.multipleParamMutex.RLock()
	defer fake.multipleParamMutex.RUnlock()
	return len(fake.multipleParamArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) MultipleParamCalls(stub func(context.Context, int32, int64, float32) ([]*multi.TypeTwo, error)) {
	fake.multipleParamMutex.Lock()
	defer fake.multipleParamMutex.Unlock()
	fake.MultipleParamStub = stub
}

func (fake *FakeTypeTwoCollectionReader) MultipleParamArgsForCall(i int) (context.Context, int32, int64, float32) {
	fake.multipleParamMutex.RLock()
	defer fake.multipleParamMutex.RUnlock()
	argsForCall := fake.multipleParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypeTwoCollectionReader) MultipleParamReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.multipleParamMutex.Lock()
	defer fake.multipleParamMutex.Unlock()
	fake.MultipleParamStub = nil
	fake.multipleParamReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) MultipleParamReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.multipleParamMutex.Lock()
	defer fake.multipleParamMutex.Unlock()
	fake.MultipleParamStub = nil
	if fake.multipleParamReturnsOnCall == nil {
		fake.multipleParamReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.multipleParamReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) OneParam(arg1 context.Context, arg2 int32) ([]*multi.TypeTwo, error) {
	fake.oneParamMutex.Lock()
	ret, specificReturn := fake.oneParamReturnsOnCall[len(fake.oneParamArgsForCall)]
	fake.oneParamArgsForCall = append(fake.oneParamArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	fake.recordInvocation("OneParam", []interface{}{arg1, arg2})
	fake.oneParamMutex.Unlock()
	if fake.OneParamStub != nil {
		return fake.OneParamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.oneParamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) OneParamCallCount() int {
	fake.oneParamMutex.RLock()
	defer fake.oneParamMutex.RUnlock()
	return len(fake.oneParamArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) OneParamCalls(stub func(context.Context, int32) ([]*multi.TypeTwo, error)) {
	fake.oneParamMutex.Lock()
	defer fake.oneParamMutex.Unlock()
	fake.OneParamStub = stub
}

func (fake *FakeTypeTwoCollectionReader) OneParamArgsForCall(i int) (context.Context, int32) {
	fake.oneParamMutex.RLock()
	defer fake.oneParamMutex.RUnlock()
	argsForCall := fake.oneParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeTwoCollectionReader) OneParamReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.oneParamMutex.Lock()
	defer fake.oneParamMutex.Unlock()
	fake.OneParamStub = nil
	fake.oneParamReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) OneParamReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.oneParamMutex.Lock()
	defer fake.oneParamMutex.Unlock()
	fake.OneParamStub = nil
	if fake.oneParamReturnsOnCall == nil {
		fake.oneParamReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.oneParamReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnly(arg1 context.Context) ([]*multi.TypeTwo, error) {
	fake.providerStubOnlyMutex.Lock()
	ret, specificReturn := fake.providerStubOnlyReturnsOnCall[len(fake.providerStubOnlyArgsForCall)]
	fake.providerStubOnlyArgsForCall = append(fake.providerStubOnlyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("ProviderStubOnly", []interface{}{arg1})
	fake.providerStubOnlyMutex.Unlock()
	if fake.ProviderStubOnlyStub != nil {
		return fake.ProviderStubOnlyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.providerStubOnlyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnlyCallCount() int {
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	return len(fake.providerStubOnlyArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnlyCalls(stub func(context.Context) ([]*multi.TypeTwo, error)) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = stub
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnlyArgsForCall(i int) context.Context {
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	argsForCall := fake.providerStubOnlyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnlyReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = nil
	fake.providerStubOnlyReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) ProviderStubOnlyReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = nil
	if fake.providerStubOnlyReturnsOnCall == nil {
		fake.providerStubOnlyReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.providerStubOnlyReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) WithComparator(arg1 context.Context, arg2 int32) ([]*multi.TypeTwo, error) {
	fake.withComparatorMutex.Lock()
	ret, specificReturn := fake.withComparatorReturnsOnCall[len(fake.withComparatorArgsForCall)]
	fake.withComparatorArgsForCall = append(fake.withComparatorArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	fake.recordInvocation("WithComparator", []interface{}{arg1, arg2})
	fake.withComparatorMutex.Unlock()
	if fake.WithComparatorStub != nil {
		return fake.WithComparatorStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.withComparatorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) WithComparatorCallCount() int {
	fake.withComparatorMutex.RLock()
	defer fake.withComparatorMutex.RUnlock()
	return len(fake.withComparatorArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) WithComparatorCalls(stub func(context.Context, int32) ([]*multi.TypeTwo, error)) {
	fake.withComparatorMutex.Lock()
	defer fake.withComparatorMutex.Unlock()
	fake.WithComparatorStub = stub
}

func (fake *FakeTypeTwoCollectionReader) WithComparatorArgsForCall(i int) (context.Context, int32) {
	fake.withComparatorMutex.RLock()
	defer fake.withComparatorMutex.RUnlock()
	argsForCall := fake.withComparatorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeTwoCollectionReader) WithComparatorReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.withComparatorMutex.Lock()
	defer fake.withComparatorMutex.Unlock()
	fake.WithComparatorStub = nil
	fake.withComparatorReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) WithComparatorReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.withComparatorMutex.Lock()
	defer fake.withComparatorMutex.Unlock()
	fake.WithComparatorStub = nil
	if fake.withComparatorReturnsOnCall == nil {
		fake.withComparatorReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.withComparatorReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) WithRest(arg1 context.Context, arg2 int32, arg3 int64, arg4 float32, arg5 float64) ([]*multi.TypeTwo, error) {
	fake.withRestMutex.Lock()
	ret, specificReturn := fake.withRestReturnsOnCall[len(fake.withRestArgsForCall)]
	fake.withRestArgsForCall = append(fake.withRestArgsForCall, struct {
		arg1 context.Context
		arg2 int32
		arg3 int64
		arg4 float32
		arg5 float64
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("WithRest", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.withRestMutex.Unlock()
	if fake.WithRestStub != nil {
		return fake.WithRestStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.withRestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeTwoCollectionReader) WithRestCallCount() int {
	fake.withRestMutex.RLock()
	defer fake.withRestMutex.RUnlock()
	return len(fake.withRestArgsForCall)
}

func (fake *FakeTypeTwoCollectionReader) WithRestCalls(stub func(context.Context, int32, int64, float32, float64) ([]*multi.TypeTwo, error)) {
	fake.withRestMutex.Lock()
	defer fake.withRestMutex.Unlock()
	fake.WithRestStub = stub
}

func (fake *FakeTypeTwoCollectionReader) WithRestArgsForCall(i int) (context.Context, int32, int64, float32, float64) {
	fake.withRestMutex.RLock()
	defer fake.withRestMutex.RUnlock()
	argsForCall := fake.withRestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTypeTwoCollectionReader) WithRestReturns(result1 []*multi.TypeTwo, result2 error) {
	fake.withRestMutex.Lock()
	defer fake.withRestMutex.Unlock()
	fake.WithRestStub = nil
	fake.withRestReturns = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) WithRestReturnsOnCall(i int, result1 []*multi.TypeTwo, result2 error) {
	fake.withRestMutex.Lock()
	defer fake.withRestMutex.Unlock()
	fake.WithRestStub = nil
	if fake.withRestReturnsOnCall == nil {
		fake.withRestReturnsOnCall = make(map[int]struct {
			result1 []*multi.TypeTwo
			result2 error
		})
	}
	fake.withRestReturnsOnCall[i] = struct {
		result1 []*multi.TypeTwo
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeTwoCollectionReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	fake.messageParamMutex.RLock()
	defer fake.messageParamMutex.RUnlock()
	fake.multipleParamMutex.RLock()
	defer fake.multipleParamMutex.RUnlock()
	fake.oneParamMutex.RLock()
	defer fake.oneParamMutex.RUnlock()
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	fake.withComparatorMutex.RLock()
	defer fake.withComparatorMutex.RUnlock()
	fake.withRestMutex.RLock()
	defer fake.withRestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypeTwoCollectionReader) recordInvocation(key string, args []interface{}) {
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

var _ dal_multi.TypeTwoCollectionReader = new(FakeTypeTwoCollectionReader)

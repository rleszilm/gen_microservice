// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"context"
	"sync"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations/types"
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/example/users"
	dal_users "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/example/users/dal"
)

type FakeUserCollectionReader struct {
	AllStub        func(context.Context) ([]*users.User, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
		arg1 context.Context
	}
	allReturns struct {
		result1 []*users.User
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	ByIdStub        func(context.Context, int64) ([]*users.User, error)
	byIdMutex       sync.RWMutex
	byIdArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	byIdReturns struct {
		result1 []*users.User
		result2 error
	}
	byIdReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	ByKindStub        func(context.Context, users.User_Kind) ([]*users.User, error)
	byKindMutex       sync.RWMutex
	byKindArgsForCall []struct {
		arg1 context.Context
		arg2 users.User_Kind
	}
	byKindReturns struct {
		result1 []*users.User
		result2 error
	}
	byKindReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	ByNameAndDivisionStub        func(context.Context, string, string) ([]*users.User, error)
	byNameAndDivisionMutex       sync.RWMutex
	byNameAndDivisionArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	byNameAndDivisionReturns struct {
		result1 []*users.User
		result2 error
	}
	byNameAndDivisionReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	ByPhoneStub        func(context.Context, *types.Phone) ([]*users.User, error)
	byPhoneMutex       sync.RWMutex
	byPhoneArgsForCall []struct {
		arg1 context.Context
		arg2 *types.Phone
	}
	byPhoneReturns struct {
		result1 []*users.User
		result2 error
	}
	byPhoneReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	FilterStub        func(context.Context, *dal_users.UserFields) ([]*users.User, error)
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 context.Context
		arg2 *dal_users.UserFields
	}
	filterReturns struct {
		result1 []*users.User
		result2 error
	}
	filterReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	InterfaceStubOnlyStub        func(context.Context) ([]*users.User, error)
	interfaceStubOnlyMutex       sync.RWMutex
	interfaceStubOnlyArgsForCall []struct {
		arg1 context.Context
	}
	interfaceStubOnlyReturns struct {
		result1 []*users.User
		result2 error
	}
	interfaceStubOnlyReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	ProviderStubOnlyStub        func(context.Context) ([]*users.User, error)
	providerStubOnlyMutex       sync.RWMutex
	providerStubOnlyArgsForCall []struct {
		arg1 context.Context
	}
	providerStubOnlyReturns struct {
		result1 []*users.User
		result2 error
	}
	providerStubOnlyReturnsOnCall map[int]struct {
		result1 []*users.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserCollectionReader) All(arg1 context.Context) ([]*users.User, error) {
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

func (fake *FakeUserCollectionReader) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *FakeUserCollectionReader) AllCalls(stub func(context.Context) ([]*users.User, error)) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = stub
}

func (fake *FakeUserCollectionReader) AllArgsForCall(i int) context.Context {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	argsForCall := fake.allArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserCollectionReader) AllReturns(result1 []*users.User, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) AllReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ById(arg1 context.Context, arg2 int64) ([]*users.User, error) {
	fake.byIdMutex.Lock()
	ret, specificReturn := fake.byIdReturnsOnCall[len(fake.byIdArgsForCall)]
	fake.byIdArgsForCall = append(fake.byIdArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("ById", []interface{}{arg1, arg2})
	fake.byIdMutex.Unlock()
	if fake.ByIdStub != nil {
		return fake.ByIdStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byIdReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionReader) ByIdCallCount() int {
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	return len(fake.byIdArgsForCall)
}

func (fake *FakeUserCollectionReader) ByIdCalls(stub func(context.Context, int64) ([]*users.User, error)) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = stub
}

func (fake *FakeUserCollectionReader) ByIdArgsForCall(i int) (context.Context, int64) {
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	argsForCall := fake.byIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionReader) ByIdReturns(result1 []*users.User, result2 error) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = nil
	fake.byIdReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByIdReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = nil
	if fake.byIdReturnsOnCall == nil {
		fake.byIdReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.byIdReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByKind(arg1 context.Context, arg2 users.User_Kind) ([]*users.User, error) {
	fake.byKindMutex.Lock()
	ret, specificReturn := fake.byKindReturnsOnCall[len(fake.byKindArgsForCall)]
	fake.byKindArgsForCall = append(fake.byKindArgsForCall, struct {
		arg1 context.Context
		arg2 users.User_Kind
	}{arg1, arg2})
	fake.recordInvocation("ByKind", []interface{}{arg1, arg2})
	fake.byKindMutex.Unlock()
	if fake.ByKindStub != nil {
		return fake.ByKindStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byKindReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionReader) ByKindCallCount() int {
	fake.byKindMutex.RLock()
	defer fake.byKindMutex.RUnlock()
	return len(fake.byKindArgsForCall)
}

func (fake *FakeUserCollectionReader) ByKindCalls(stub func(context.Context, users.User_Kind) ([]*users.User, error)) {
	fake.byKindMutex.Lock()
	defer fake.byKindMutex.Unlock()
	fake.ByKindStub = stub
}

func (fake *FakeUserCollectionReader) ByKindArgsForCall(i int) (context.Context, users.User_Kind) {
	fake.byKindMutex.RLock()
	defer fake.byKindMutex.RUnlock()
	argsForCall := fake.byKindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionReader) ByKindReturns(result1 []*users.User, result2 error) {
	fake.byKindMutex.Lock()
	defer fake.byKindMutex.Unlock()
	fake.ByKindStub = nil
	fake.byKindReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByKindReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.byKindMutex.Lock()
	defer fake.byKindMutex.Unlock()
	fake.ByKindStub = nil
	if fake.byKindReturnsOnCall == nil {
		fake.byKindReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.byKindReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByNameAndDivision(arg1 context.Context, arg2 string, arg3 string) ([]*users.User, error) {
	fake.byNameAndDivisionMutex.Lock()
	ret, specificReturn := fake.byNameAndDivisionReturnsOnCall[len(fake.byNameAndDivisionArgsForCall)]
	fake.byNameAndDivisionArgsForCall = append(fake.byNameAndDivisionArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("ByNameAndDivision", []interface{}{arg1, arg2, arg3})
	fake.byNameAndDivisionMutex.Unlock()
	if fake.ByNameAndDivisionStub != nil {
		return fake.ByNameAndDivisionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byNameAndDivisionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionReader) ByNameAndDivisionCallCount() int {
	fake.byNameAndDivisionMutex.RLock()
	defer fake.byNameAndDivisionMutex.RUnlock()
	return len(fake.byNameAndDivisionArgsForCall)
}

func (fake *FakeUserCollectionReader) ByNameAndDivisionCalls(stub func(context.Context, string, string) ([]*users.User, error)) {
	fake.byNameAndDivisionMutex.Lock()
	defer fake.byNameAndDivisionMutex.Unlock()
	fake.ByNameAndDivisionStub = stub
}

func (fake *FakeUserCollectionReader) ByNameAndDivisionArgsForCall(i int) (context.Context, string, string) {
	fake.byNameAndDivisionMutex.RLock()
	defer fake.byNameAndDivisionMutex.RUnlock()
	argsForCall := fake.byNameAndDivisionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserCollectionReader) ByNameAndDivisionReturns(result1 []*users.User, result2 error) {
	fake.byNameAndDivisionMutex.Lock()
	defer fake.byNameAndDivisionMutex.Unlock()
	fake.ByNameAndDivisionStub = nil
	fake.byNameAndDivisionReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByNameAndDivisionReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.byNameAndDivisionMutex.Lock()
	defer fake.byNameAndDivisionMutex.Unlock()
	fake.ByNameAndDivisionStub = nil
	if fake.byNameAndDivisionReturnsOnCall == nil {
		fake.byNameAndDivisionReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.byNameAndDivisionReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByPhone(arg1 context.Context, arg2 *types.Phone) ([]*users.User, error) {
	fake.byPhoneMutex.Lock()
	ret, specificReturn := fake.byPhoneReturnsOnCall[len(fake.byPhoneArgsForCall)]
	fake.byPhoneArgsForCall = append(fake.byPhoneArgsForCall, struct {
		arg1 context.Context
		arg2 *types.Phone
	}{arg1, arg2})
	fake.recordInvocation("ByPhone", []interface{}{arg1, arg2})
	fake.byPhoneMutex.Unlock()
	if fake.ByPhoneStub != nil {
		return fake.ByPhoneStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byPhoneReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserCollectionReader) ByPhoneCallCount() int {
	fake.byPhoneMutex.RLock()
	defer fake.byPhoneMutex.RUnlock()
	return len(fake.byPhoneArgsForCall)
}

func (fake *FakeUserCollectionReader) ByPhoneCalls(stub func(context.Context, *types.Phone) ([]*users.User, error)) {
	fake.byPhoneMutex.Lock()
	defer fake.byPhoneMutex.Unlock()
	fake.ByPhoneStub = stub
}

func (fake *FakeUserCollectionReader) ByPhoneArgsForCall(i int) (context.Context, *types.Phone) {
	fake.byPhoneMutex.RLock()
	defer fake.byPhoneMutex.RUnlock()
	argsForCall := fake.byPhoneArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionReader) ByPhoneReturns(result1 []*users.User, result2 error) {
	fake.byPhoneMutex.Lock()
	defer fake.byPhoneMutex.Unlock()
	fake.ByPhoneStub = nil
	fake.byPhoneReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ByPhoneReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.byPhoneMutex.Lock()
	defer fake.byPhoneMutex.Unlock()
	fake.ByPhoneStub = nil
	if fake.byPhoneReturnsOnCall == nil {
		fake.byPhoneReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.byPhoneReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) Filter(arg1 context.Context, arg2 *dal_users.UserFields) ([]*users.User, error) {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 context.Context
		arg2 *dal_users.UserFields
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

func (fake *FakeUserCollectionReader) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *FakeUserCollectionReader) FilterCalls(stub func(context.Context, *dal_users.UserFields) ([]*users.User, error)) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = stub
}

func (fake *FakeUserCollectionReader) FilterArgsForCall(i int) (context.Context, *dal_users.UserFields) {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	argsForCall := fake.filterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserCollectionReader) FilterReturns(result1 []*users.User, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) FilterReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) InterfaceStubOnly(arg1 context.Context) ([]*users.User, error) {
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

func (fake *FakeUserCollectionReader) InterfaceStubOnlyCallCount() int {
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	return len(fake.interfaceStubOnlyArgsForCall)
}

func (fake *FakeUserCollectionReader) InterfaceStubOnlyCalls(stub func(context.Context) ([]*users.User, error)) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = stub
}

func (fake *FakeUserCollectionReader) InterfaceStubOnlyArgsForCall(i int) context.Context {
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	argsForCall := fake.interfaceStubOnlyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserCollectionReader) InterfaceStubOnlyReturns(result1 []*users.User, result2 error) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = nil
	fake.interfaceStubOnlyReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) InterfaceStubOnlyReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.interfaceStubOnlyMutex.Lock()
	defer fake.interfaceStubOnlyMutex.Unlock()
	fake.InterfaceStubOnlyStub = nil
	if fake.interfaceStubOnlyReturnsOnCall == nil {
		fake.interfaceStubOnlyReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.interfaceStubOnlyReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ProviderStubOnly(arg1 context.Context) ([]*users.User, error) {
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

func (fake *FakeUserCollectionReader) ProviderStubOnlyCallCount() int {
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	return len(fake.providerStubOnlyArgsForCall)
}

func (fake *FakeUserCollectionReader) ProviderStubOnlyCalls(stub func(context.Context) ([]*users.User, error)) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = stub
}

func (fake *FakeUserCollectionReader) ProviderStubOnlyArgsForCall(i int) context.Context {
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	argsForCall := fake.providerStubOnlyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserCollectionReader) ProviderStubOnlyReturns(result1 []*users.User, result2 error) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = nil
	fake.providerStubOnlyReturns = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) ProviderStubOnlyReturnsOnCall(i int, result1 []*users.User, result2 error) {
	fake.providerStubOnlyMutex.Lock()
	defer fake.providerStubOnlyMutex.Unlock()
	fake.ProviderStubOnlyStub = nil
	if fake.providerStubOnlyReturnsOnCall == nil {
		fake.providerStubOnlyReturnsOnCall = make(map[int]struct {
			result1 []*users.User
			result2 error
		})
	}
	fake.providerStubOnlyReturnsOnCall[i] = struct {
		result1 []*users.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserCollectionReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	fake.byKindMutex.RLock()
	defer fake.byKindMutex.RUnlock()
	fake.byNameAndDivisionMutex.RLock()
	defer fake.byNameAndDivisionMutex.RUnlock()
	fake.byPhoneMutex.RLock()
	defer fake.byPhoneMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	fake.interfaceStubOnlyMutex.RLock()
	defer fake.interfaceStubOnlyMutex.RUnlock()
	fake.providerStubOnlyMutex.RLock()
	defer fake.providerStubOnlyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserCollectionReader) recordInvocation(key string, args []interface{}) {
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

var _ dal_users.UserCollectionReader = new(FakeUserCollectionReader)

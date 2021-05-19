// Code generated by counterfeiter. DO NOT EDIT.
package repfakes

import (
	"sync"

	"code.cloudfoundry.org/rep"
)

type FakeClientFactory struct {
	CreateClientStub        func(string, string) (rep.Client, error)
	createClientMutex       sync.RWMutex
	createClientArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createClientReturns struct {
		result1 rep.Client
		result2 error
	}
	createClientReturnsOnCall map[int]struct {
		result1 rep.Client
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientFactory) CreateClient(arg1 string, arg2 string) (rep.Client, error) {
	fake.createClientMutex.Lock()
	ret, specificReturn := fake.createClientReturnsOnCall[len(fake.createClientArgsForCall)]
	fake.createClientArgsForCall = append(fake.createClientArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateClient", []interface{}{arg1, arg2})
	createClientStubCopy := fake.CreateClientStub
	fake.createClientMutex.Unlock()
	if createClientStubCopy != nil {
		return createClientStubCopy(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientFactory) CreateClientCallCount() int {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return len(fake.createClientArgsForCall)
}

func (fake *FakeClientFactory) CreateClientCalls(stub func(string, string) (rep.Client, error)) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = stub
}

func (fake *FakeClientFactory) CreateClientArgsForCall(i int) (string, string) {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	argsForCall := fake.createClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClientFactory) CreateClientReturns(result1 rep.Client, result2 error) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = nil
	fake.createClientReturns = struct {
		result1 rep.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) CreateClientReturnsOnCall(i int, result1 rep.Client, result2 error) {
	fake.createClientMutex.Lock()
	defer fake.createClientMutex.Unlock()
	fake.CreateClientStub = nil
	if fake.createClientReturnsOnCall == nil {
		fake.createClientReturnsOnCall = make(map[int]struct {
			result1 rep.Client
			result2 error
		})
	}
	fake.createClientReturnsOnCall[i] = struct {
		result1 rep.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ rep.ClientFactory = new(FakeClientFactory)
// Code generated by counterfeiter. DO NOT EDIT.
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeCreatingVolume struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	CreatedStub        func() (dbng.CreatedVolume, error)
	createdMutex       sync.RWMutex
	createdArgsForCall []struct{}
	createdReturns     struct {
		result1 dbng.CreatedVolume
		result2 error
	}
	createdReturnsOnCall map[int]struct {
		result1 dbng.CreatedVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreatingVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handleReturns.result1
}

func (fake *FakeCreatingVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeCreatingVolume) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatingVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatingVolume) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeCreatingVolume) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeCreatingVolume) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatingVolume) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatingVolume) Created() (dbng.CreatedVolume, error) {
	fake.createdMutex.Lock()
	ret, specificReturn := fake.createdReturnsOnCall[len(fake.createdArgsForCall)]
	fake.createdArgsForCall = append(fake.createdArgsForCall, struct{}{})
	fake.recordInvocation("Created", []interface{}{})
	fake.createdMutex.Unlock()
	if fake.CreatedStub != nil {
		return fake.CreatedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createdReturns.result1, fake.createdReturns.result2
}

func (fake *FakeCreatingVolume) CreatedCallCount() int {
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	return len(fake.createdArgsForCall)
}

func (fake *FakeCreatingVolume) CreatedReturns(result1 dbng.CreatedVolume, result2 error) {
	fake.CreatedStub = nil
	fake.createdReturns = struct {
		result1 dbng.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatingVolume) CreatedReturnsOnCall(i int, result1 dbng.CreatedVolume, result2 error) {
	fake.CreatedStub = nil
	if fake.createdReturnsOnCall == nil {
		fake.createdReturnsOnCall = make(map[int]struct {
			result1 dbng.CreatedVolume
			result2 error
		})
	}
	fake.createdReturnsOnCall[i] = struct {
		result1 dbng.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatingVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreatingVolume) recordInvocation(key string, args []interface{}) {
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

var _ dbng.CreatingVolume = new(FakeCreatingVolume)

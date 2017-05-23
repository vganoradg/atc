// Code generated by counterfeiter. DO NOT EDIT.
package cessnafakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/atc/cessna"
	"github.com/concourse/baggageclaim"
)

type FakeWorker struct {
	GardenClientStub        func() garden.Client
	gardenClientMutex       sync.RWMutex
	gardenClientArgsForCall []struct{}
	gardenClientReturns     struct {
		result1 garden.Client
	}
	gardenClientReturnsOnCall map[int]struct {
		result1 garden.Client
	}
	BaggageClaimClientStub        func() baggageclaim.Client
	baggageClaimClientMutex       sync.RWMutex
	baggageClaimClientArgsForCall []struct{}
	baggageClaimClientReturns     struct {
		result1 baggageclaim.Client
	}
	baggageClaimClientReturnsOnCall map[int]struct {
		result1 baggageclaim.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) GardenClient() garden.Client {
	fake.gardenClientMutex.Lock()
	ret, specificReturn := fake.gardenClientReturnsOnCall[len(fake.gardenClientArgsForCall)]
	fake.gardenClientArgsForCall = append(fake.gardenClientArgsForCall, struct{}{})
	fake.recordInvocation("GardenClient", []interface{}{})
	fake.gardenClientMutex.Unlock()
	if fake.GardenClientStub != nil {
		return fake.GardenClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gardenClientReturns.result1
}

func (fake *FakeWorker) GardenClientCallCount() int {
	fake.gardenClientMutex.RLock()
	defer fake.gardenClientMutex.RUnlock()
	return len(fake.gardenClientArgsForCall)
}

func (fake *FakeWorker) GardenClientReturns(result1 garden.Client) {
	fake.GardenClientStub = nil
	fake.gardenClientReturns = struct {
		result1 garden.Client
	}{result1}
}

func (fake *FakeWorker) GardenClientReturnsOnCall(i int, result1 garden.Client) {
	fake.GardenClientStub = nil
	if fake.gardenClientReturnsOnCall == nil {
		fake.gardenClientReturnsOnCall = make(map[int]struct {
			result1 garden.Client
		})
	}
	fake.gardenClientReturnsOnCall[i] = struct {
		result1 garden.Client
	}{result1}
}

func (fake *FakeWorker) BaggageClaimClient() baggageclaim.Client {
	fake.baggageClaimClientMutex.Lock()
	ret, specificReturn := fake.baggageClaimClientReturnsOnCall[len(fake.baggageClaimClientArgsForCall)]
	fake.baggageClaimClientArgsForCall = append(fake.baggageClaimClientArgsForCall, struct{}{})
	fake.recordInvocation("BaggageClaimClient", []interface{}{})
	fake.baggageClaimClientMutex.Unlock()
	if fake.BaggageClaimClientStub != nil {
		return fake.BaggageClaimClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.baggageClaimClientReturns.result1
}

func (fake *FakeWorker) BaggageClaimClientCallCount() int {
	fake.baggageClaimClientMutex.RLock()
	defer fake.baggageClaimClientMutex.RUnlock()
	return len(fake.baggageClaimClientArgsForCall)
}

func (fake *FakeWorker) BaggageClaimClientReturns(result1 baggageclaim.Client) {
	fake.BaggageClaimClientStub = nil
	fake.baggageClaimClientReturns = struct {
		result1 baggageclaim.Client
	}{result1}
}

func (fake *FakeWorker) BaggageClaimClientReturnsOnCall(i int, result1 baggageclaim.Client) {
	fake.BaggageClaimClientStub = nil
	if fake.baggageClaimClientReturnsOnCall == nil {
		fake.baggageClaimClientReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Client
		})
	}
	fake.baggageClaimClientReturnsOnCall[i] = struct {
		result1 baggageclaim.Client
	}{result1}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gardenClientMutex.RLock()
	defer fake.gardenClientMutex.RUnlock()
	fake.baggageClaimClientMutex.RLock()
	defer fake.baggageClaimClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
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

var _ cessna.Worker = new(FakeWorker)

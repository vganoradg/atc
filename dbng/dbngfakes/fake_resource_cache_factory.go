// Code generated by counterfeiter. DO NOT EDIT.
package dbngfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeResourceCacheFactory struct {
	FindOrCreateResourceCacheStub        func(logger lager.Logger, resourceUser dbng.ResourceUser, resourceTypeName string, version atc.Version, source atc.Source, params atc.Params, resourceTypes atc.VersionedResourceTypes) (*dbng.UsedResourceCache, error)
	findOrCreateResourceCacheMutex       sync.RWMutex
	findOrCreateResourceCacheArgsForCall []struct {
		logger           lager.Logger
		resourceUser     dbng.ResourceUser
		resourceTypeName string
		version          atc.Version
		source           atc.Source
		params           atc.Params
		resourceTypes    atc.VersionedResourceTypes
	}
	findOrCreateResourceCacheReturns struct {
		result1 *dbng.UsedResourceCache
		result2 error
	}
	findOrCreateResourceCacheReturnsOnCall map[int]struct {
		result1 *dbng.UsedResourceCache
		result2 error
	}
	CleanUsesForFinishedBuildsStub        func() error
	cleanUsesForFinishedBuildsMutex       sync.RWMutex
	cleanUsesForFinishedBuildsArgsForCall []struct{}
	cleanUsesForFinishedBuildsReturns     struct {
		result1 error
	}
	cleanUsesForFinishedBuildsReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUsesForInactiveResourceTypesStub        func() error
	cleanUsesForInactiveResourceTypesMutex       sync.RWMutex
	cleanUsesForInactiveResourceTypesArgsForCall []struct{}
	cleanUsesForInactiveResourceTypesReturns     struct {
		result1 error
	}
	cleanUsesForInactiveResourceTypesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUsesForInactiveResourcesStub        func() error
	cleanUsesForInactiveResourcesMutex       sync.RWMutex
	cleanUsesForInactiveResourcesArgsForCall []struct{}
	cleanUsesForInactiveResourcesReturns     struct {
		result1 error
	}
	cleanUsesForInactiveResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUsesForPausedPipelineResourcesStub        func() error
	cleanUsesForPausedPipelineResourcesMutex       sync.RWMutex
	cleanUsesForPausedPipelineResourcesArgsForCall []struct{}
	cleanUsesForPausedPipelineResourcesReturns     struct {
		result1 error
	}
	cleanUsesForPausedPipelineResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUpInvalidCachesStub        func() error
	cleanUpInvalidCachesMutex       sync.RWMutex
	cleanUpInvalidCachesArgsForCall []struct{}
	cleanUpInvalidCachesReturns     struct {
		result1 error
	}
	cleanUpInvalidCachesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCache(logger lager.Logger, resourceUser dbng.ResourceUser, resourceTypeName string, version atc.Version, source atc.Source, params atc.Params, resourceTypes atc.VersionedResourceTypes) (*dbng.UsedResourceCache, error) {
	fake.findOrCreateResourceCacheMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceCacheReturnsOnCall[len(fake.findOrCreateResourceCacheArgsForCall)]
	fake.findOrCreateResourceCacheArgsForCall = append(fake.findOrCreateResourceCacheArgsForCall, struct {
		logger           lager.Logger
		resourceUser     dbng.ResourceUser
		resourceTypeName string
		version          atc.Version
		source           atc.Source
		params           atc.Params
		resourceTypes    atc.VersionedResourceTypes
	}{logger, resourceUser, resourceTypeName, version, source, params, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceCache", []interface{}{logger, resourceUser, resourceTypeName, version, source, params, resourceTypes})
	fake.findOrCreateResourceCacheMutex.Unlock()
	if fake.FindOrCreateResourceCacheStub != nil {
		return fake.FindOrCreateResourceCacheStub(logger, resourceUser, resourceTypeName, version, source, params, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateResourceCacheReturns.result1, fake.findOrCreateResourceCacheReturns.result2
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheCallCount() int {
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	return len(fake.findOrCreateResourceCacheArgsForCall)
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheArgsForCall(i int) (lager.Logger, dbng.ResourceUser, string, atc.Version, atc.Source, atc.Params, atc.VersionedResourceTypes) {
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	return fake.findOrCreateResourceCacheArgsForCall[i].logger, fake.findOrCreateResourceCacheArgsForCall[i].resourceUser, fake.findOrCreateResourceCacheArgsForCall[i].resourceTypeName, fake.findOrCreateResourceCacheArgsForCall[i].version, fake.findOrCreateResourceCacheArgsForCall[i].source, fake.findOrCreateResourceCacheArgsForCall[i].params, fake.findOrCreateResourceCacheArgsForCall[i].resourceTypes
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheReturns(result1 *dbng.UsedResourceCache, result2 error) {
	fake.FindOrCreateResourceCacheStub = nil
	fake.findOrCreateResourceCacheReturns = struct {
		result1 *dbng.UsedResourceCache
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheReturnsOnCall(i int, result1 *dbng.UsedResourceCache, result2 error) {
	fake.FindOrCreateResourceCacheStub = nil
	if fake.findOrCreateResourceCacheReturnsOnCall == nil {
		fake.findOrCreateResourceCacheReturnsOnCall = make(map[int]struct {
			result1 *dbng.UsedResourceCache
			result2 error
		})
	}
	fake.findOrCreateResourceCacheReturnsOnCall[i] = struct {
		result1 *dbng.UsedResourceCache
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) CleanUsesForFinishedBuilds() error {
	fake.cleanUsesForFinishedBuildsMutex.Lock()
	ret, specificReturn := fake.cleanUsesForFinishedBuildsReturnsOnCall[len(fake.cleanUsesForFinishedBuildsArgsForCall)]
	fake.cleanUsesForFinishedBuildsArgsForCall = append(fake.cleanUsesForFinishedBuildsArgsForCall, struct{}{})
	fake.recordInvocation("CleanUsesForFinishedBuilds", []interface{}{})
	fake.cleanUsesForFinishedBuildsMutex.Unlock()
	if fake.CleanUsesForFinishedBuildsStub != nil {
		return fake.CleanUsesForFinishedBuildsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUsesForFinishedBuildsReturns.result1
}

func (fake *FakeResourceCacheFactory) CleanUsesForFinishedBuildsCallCount() int {
	fake.cleanUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanUsesForFinishedBuildsMutex.RUnlock()
	return len(fake.cleanUsesForFinishedBuildsArgsForCall)
}

func (fake *FakeResourceCacheFactory) CleanUsesForFinishedBuildsReturns(result1 error) {
	fake.CleanUsesForFinishedBuildsStub = nil
	fake.cleanUsesForFinishedBuildsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForFinishedBuildsReturnsOnCall(i int, result1 error) {
	fake.CleanUsesForFinishedBuildsStub = nil
	if fake.cleanUsesForFinishedBuildsReturnsOnCall == nil {
		fake.cleanUsesForFinishedBuildsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUsesForFinishedBuildsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourceTypes() error {
	fake.cleanUsesForInactiveResourceTypesMutex.Lock()
	ret, specificReturn := fake.cleanUsesForInactiveResourceTypesReturnsOnCall[len(fake.cleanUsesForInactiveResourceTypesArgsForCall)]
	fake.cleanUsesForInactiveResourceTypesArgsForCall = append(fake.cleanUsesForInactiveResourceTypesArgsForCall, struct{}{})
	fake.recordInvocation("CleanUsesForInactiveResourceTypes", []interface{}{})
	fake.cleanUsesForInactiveResourceTypesMutex.Unlock()
	if fake.CleanUsesForInactiveResourceTypesStub != nil {
		return fake.CleanUsesForInactiveResourceTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUsesForInactiveResourceTypesReturns.result1
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourceTypesCallCount() int {
	fake.cleanUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanUsesForInactiveResourceTypesMutex.RUnlock()
	return len(fake.cleanUsesForInactiveResourceTypesArgsForCall)
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourceTypesReturns(result1 error) {
	fake.CleanUsesForInactiveResourceTypesStub = nil
	fake.cleanUsesForInactiveResourceTypesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourceTypesReturnsOnCall(i int, result1 error) {
	fake.CleanUsesForInactiveResourceTypesStub = nil
	if fake.cleanUsesForInactiveResourceTypesReturnsOnCall == nil {
		fake.cleanUsesForInactiveResourceTypesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUsesForInactiveResourceTypesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResources() error {
	fake.cleanUsesForInactiveResourcesMutex.Lock()
	ret, specificReturn := fake.cleanUsesForInactiveResourcesReturnsOnCall[len(fake.cleanUsesForInactiveResourcesArgsForCall)]
	fake.cleanUsesForInactiveResourcesArgsForCall = append(fake.cleanUsesForInactiveResourcesArgsForCall, struct{}{})
	fake.recordInvocation("CleanUsesForInactiveResources", []interface{}{})
	fake.cleanUsesForInactiveResourcesMutex.Unlock()
	if fake.CleanUsesForInactiveResourcesStub != nil {
		return fake.CleanUsesForInactiveResourcesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUsesForInactiveResourcesReturns.result1
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourcesCallCount() int {
	fake.cleanUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanUsesForInactiveResourcesMutex.RUnlock()
	return len(fake.cleanUsesForInactiveResourcesArgsForCall)
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourcesReturns(result1 error) {
	fake.CleanUsesForInactiveResourcesStub = nil
	fake.cleanUsesForInactiveResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForInactiveResourcesReturnsOnCall(i int, result1 error) {
	fake.CleanUsesForInactiveResourcesStub = nil
	if fake.cleanUsesForInactiveResourcesReturnsOnCall == nil {
		fake.cleanUsesForInactiveResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUsesForInactiveResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForPausedPipelineResources() error {
	fake.cleanUsesForPausedPipelineResourcesMutex.Lock()
	ret, specificReturn := fake.cleanUsesForPausedPipelineResourcesReturnsOnCall[len(fake.cleanUsesForPausedPipelineResourcesArgsForCall)]
	fake.cleanUsesForPausedPipelineResourcesArgsForCall = append(fake.cleanUsesForPausedPipelineResourcesArgsForCall, struct{}{})
	fake.recordInvocation("CleanUsesForPausedPipelineResources", []interface{}{})
	fake.cleanUsesForPausedPipelineResourcesMutex.Unlock()
	if fake.CleanUsesForPausedPipelineResourcesStub != nil {
		return fake.CleanUsesForPausedPipelineResourcesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUsesForPausedPipelineResourcesReturns.result1
}

func (fake *FakeResourceCacheFactory) CleanUsesForPausedPipelineResourcesCallCount() int {
	fake.cleanUsesForPausedPipelineResourcesMutex.RLock()
	defer fake.cleanUsesForPausedPipelineResourcesMutex.RUnlock()
	return len(fake.cleanUsesForPausedPipelineResourcesArgsForCall)
}

func (fake *FakeResourceCacheFactory) CleanUsesForPausedPipelineResourcesReturns(result1 error) {
	fake.CleanUsesForPausedPipelineResourcesStub = nil
	fake.cleanUsesForPausedPipelineResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUsesForPausedPipelineResourcesReturnsOnCall(i int, result1 error) {
	fake.CleanUsesForPausedPipelineResourcesStub = nil
	if fake.cleanUsesForPausedPipelineResourcesReturnsOnCall == nil {
		fake.cleanUsesForPausedPipelineResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUsesForPausedPipelineResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUpInvalidCaches() error {
	fake.cleanUpInvalidCachesMutex.Lock()
	ret, specificReturn := fake.cleanUpInvalidCachesReturnsOnCall[len(fake.cleanUpInvalidCachesArgsForCall)]
	fake.cleanUpInvalidCachesArgsForCall = append(fake.cleanUpInvalidCachesArgsForCall, struct{}{})
	fake.recordInvocation("CleanUpInvalidCaches", []interface{}{})
	fake.cleanUpInvalidCachesMutex.Unlock()
	if fake.CleanUpInvalidCachesStub != nil {
		return fake.CleanUpInvalidCachesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUpInvalidCachesReturns.result1
}

func (fake *FakeResourceCacheFactory) CleanUpInvalidCachesCallCount() int {
	fake.cleanUpInvalidCachesMutex.RLock()
	defer fake.cleanUpInvalidCachesMutex.RUnlock()
	return len(fake.cleanUpInvalidCachesArgsForCall)
}

func (fake *FakeResourceCacheFactory) CleanUpInvalidCachesReturns(result1 error) {
	fake.CleanUpInvalidCachesStub = nil
	fake.cleanUpInvalidCachesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) CleanUpInvalidCachesReturnsOnCall(i int, result1 error) {
	fake.CleanUpInvalidCachesStub = nil
	if fake.cleanUpInvalidCachesReturnsOnCall == nil {
		fake.cleanUpInvalidCachesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUpInvalidCachesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	fake.cleanUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanUsesForFinishedBuildsMutex.RUnlock()
	fake.cleanUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanUsesForInactiveResourceTypesMutex.RUnlock()
	fake.cleanUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanUsesForInactiveResourcesMutex.RUnlock()
	fake.cleanUsesForPausedPipelineResourcesMutex.RLock()
	defer fake.cleanUsesForPausedPipelineResourcesMutex.RUnlock()
	fake.cleanUpInvalidCachesMutex.RLock()
	defer fake.cleanUpInvalidCachesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceCacheFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.ResourceCacheFactory = new(FakeResourceCacheFactory)

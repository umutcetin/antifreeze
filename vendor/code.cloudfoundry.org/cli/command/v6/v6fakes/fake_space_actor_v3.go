// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeSpaceActorV3 struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetEffectiveIsolationSegmentBySpaceStub        func(string, string) (v3action.IsolationSegment, v3action.Warnings, error)
	getEffectiveIsolationSegmentBySpaceMutex       sync.RWMutex
	getEffectiveIsolationSegmentBySpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getEffectiveIsolationSegmentBySpaceReturns struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}
	getEffectiveIsolationSegmentBySpaceReturnsOnCall map[int]struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpace(arg1 string, arg2 string) (v3action.IsolationSegment, v3action.Warnings, error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	ret, specificReturn := fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)]
	fake.getEffectiveIsolationSegmentBySpaceArgsForCall = append(fake.getEffectiveIsolationSegmentBySpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetEffectiveIsolationSegmentBySpace", []interface{}{arg1, arg2})
	fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	if fake.GetEffectiveIsolationSegmentBySpaceStub != nil {
		return fake.GetEffectiveIsolationSegmentBySpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getEffectiveIsolationSegmentBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceCallCount() int {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	return len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceCalls(stub func(string, string) (v3action.IsolationSegment, v3action.Warnings, error)) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = stub
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceArgsForCall(i int) (string, string) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	argsForCall := fake.getEffectiveIsolationSegmentBySpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceReturns(result1 v3action.IsolationSegment, result2 v3action.Warnings, result3 error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	fake.getEffectiveIsolationSegmentBySpaceReturns = struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceReturnsOnCall(i int, result1 v3action.IsolationSegment, result2 v3action.Warnings, result3 error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	if fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall == nil {
		fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.IsolationSegment
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[i] = struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActorV3) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceActorV3) recordInvocation(key string, args []interface{}) {
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

var _ v6.SpaceActorV3 = new(FakeSpaceActorV3)

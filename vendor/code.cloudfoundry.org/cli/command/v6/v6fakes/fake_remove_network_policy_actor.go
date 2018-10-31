// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	cfnetworkingaction "code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeRemoveNetworkPolicyActor struct {
	RemoveNetworkPolicyStub        func(string, string, string, string, int, int) (cfnetworkingaction.Warnings, error)
	removeNetworkPolicyMutex       sync.RWMutex
	removeNetworkPolicyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 int
		arg6 int
	}
	removeNetworkPolicyReturns struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}
	removeNetworkPolicyReturnsOnCall map[int]struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicy(arg1 string, arg2 string, arg3 string, arg4 string, arg5 int, arg6 int) (cfnetworkingaction.Warnings, error) {
	fake.removeNetworkPolicyMutex.Lock()
	ret, specificReturn := fake.removeNetworkPolicyReturnsOnCall[len(fake.removeNetworkPolicyArgsForCall)]
	fake.removeNetworkPolicyArgsForCall = append(fake.removeNetworkPolicyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 int
		arg6 int
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("RemoveNetworkPolicy", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.removeNetworkPolicyMutex.Unlock()
	if fake.RemoveNetworkPolicyStub != nil {
		return fake.RemoveNetworkPolicyStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.removeNetworkPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicyCallCount() int {
	fake.removeNetworkPolicyMutex.RLock()
	defer fake.removeNetworkPolicyMutex.RUnlock()
	return len(fake.removeNetworkPolicyArgsForCall)
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicyCalls(stub func(string, string, string, string, int, int) (cfnetworkingaction.Warnings, error)) {
	fake.removeNetworkPolicyMutex.Lock()
	defer fake.removeNetworkPolicyMutex.Unlock()
	fake.RemoveNetworkPolicyStub = stub
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicyArgsForCall(i int) (string, string, string, string, int, int) {
	fake.removeNetworkPolicyMutex.RLock()
	defer fake.removeNetworkPolicyMutex.RUnlock()
	argsForCall := fake.removeNetworkPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicyReturns(result1 cfnetworkingaction.Warnings, result2 error) {
	fake.removeNetworkPolicyMutex.Lock()
	defer fake.removeNetworkPolicyMutex.Unlock()
	fake.RemoveNetworkPolicyStub = nil
	fake.removeNetworkPolicyReturns = struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoveNetworkPolicyActor) RemoveNetworkPolicyReturnsOnCall(i int, result1 cfnetworkingaction.Warnings, result2 error) {
	fake.removeNetworkPolicyMutex.Lock()
	defer fake.removeNetworkPolicyMutex.Unlock()
	fake.RemoveNetworkPolicyStub = nil
	if fake.removeNetworkPolicyReturnsOnCall == nil {
		fake.removeNetworkPolicyReturnsOnCall = make(map[int]struct {
			result1 cfnetworkingaction.Warnings
			result2 error
		})
	}
	fake.removeNetworkPolicyReturnsOnCall[i] = struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoveNetworkPolicyActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeNetworkPolicyMutex.RLock()
	defer fake.removeNetworkPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRemoveNetworkPolicyActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.RemoveNetworkPolicyActor = new(FakeRemoveNetworkPolicyActor)

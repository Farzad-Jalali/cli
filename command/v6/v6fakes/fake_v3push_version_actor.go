// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cli/actor/loggingaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3PushVersionActor struct {
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, loggingaction.LogCacheClient) (<-chan loggingaction.LogMessage, <-chan error, v3action.Warnings, error, context.CancelFunc)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 loggingaction.LogCacheClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan loggingaction.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
		result5 context.CancelFunc
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan loggingaction.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
		result5 context.CancelFunc
	}
	PollStartStub        func(string, chan<- v3action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	RestartApplicationStub        func(string) (v3action.Warnings, error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		arg1 string
	}
	restartApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 loggingaction.LogCacheClient) (<-chan loggingaction.LogMessage, <-chan error, v3action.Warnings, error, context.CancelFunc) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 loggingaction.LogCacheClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4, fakeReturns.result5
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, loggingaction.LogCacheClient) (<-chan loggingaction.LogMessage, <-chan error, v3action.Warnings, error, context.CancelFunc)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, loggingaction.LogCacheClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan loggingaction.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error, result5 context.CancelFunc) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan loggingaction.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
		result5 context.CancelFunc
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan loggingaction.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error, result5 context.CancelFunc) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan loggingaction.LogMessage
			result2 <-chan error
			result3 v3action.Warnings
			result4 error
			result5 context.CancelFunc
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan loggingaction.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
		result5 context.CancelFunc
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeV3PushVersionActor) PollStart(arg1 string, arg2 chan<- v3action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}{arg1, arg2})
	fake.recordInvocation("PollStart", []interface{}{arg1, arg2})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pollStartReturns
	return fakeReturns.result1
}

func (fake *FakeV3PushVersionActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeV3PushVersionActor) PollStartCalls(stub func(string, chan<- v3action.Warnings) error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = stub
}

func (fake *FakeV3PushVersionActor) PollStartArgsForCall(i int) (string, chan<- v3action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	argsForCall := fake.pollStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3PushVersionActor) PollStartReturns(result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushVersionActor) PollStartReturnsOnCall(i int, result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushVersionActor) RestartApplication(arg1 string) (v3action.Warnings, error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RestartApplication", []interface{}{arg1})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.restartApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV3PushVersionActor) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakeV3PushVersionActor) RestartApplicationCalls(stub func(string) (v3action.Warnings, error)) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = stub
}

func (fake *FakeV3PushVersionActor) RestartApplicationArgsForCall(i int) string {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	argsForCall := fake.restartApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3PushVersionActor) RestartApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushVersionActor) RestartApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushVersionActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3PushVersionActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3PushVersionActor = new(FakeV3PushVersionActor)

// Code generated by counterfeiter. DO NOT EDIT.
package wrapperfakes

import (
	sync "sync"
	time "time"

	wrapper "code.cloudfoundry.org/cli/api/cloudcontroller/wrapper"
)

type FakeRequestLoggerOutput struct {
	DisplayHeaderStub        func(string, string) error
	displayHeaderMutex       sync.RWMutex
	displayHeaderArgsForCall []struct {
		arg1 string
		arg2 string
	}
	displayHeaderReturns struct {
		result1 error
	}
	displayHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayHostStub        func(string) error
	displayHostMutex       sync.RWMutex
	displayHostArgsForCall []struct {
		arg1 string
	}
	displayHostReturns struct {
		result1 error
	}
	displayHostReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayJSONBodyStub        func([]byte) error
	displayJSONBodyMutex       sync.RWMutex
	displayJSONBodyArgsForCall []struct {
		arg1 []byte
	}
	displayJSONBodyReturns struct {
		result1 error
	}
	displayJSONBodyReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayMessageStub        func(string) error
	displayMessageMutex       sync.RWMutex
	displayMessageArgsForCall []struct {
		arg1 string
	}
	displayMessageReturns struct {
		result1 error
	}
	displayMessageReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayRequestHeaderStub        func(string, string, string) error
	displayRequestHeaderMutex       sync.RWMutex
	displayRequestHeaderArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	displayRequestHeaderReturns struct {
		result1 error
	}
	displayRequestHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayResponseHeaderStub        func(string, string) error
	displayResponseHeaderMutex       sync.RWMutex
	displayResponseHeaderArgsForCall []struct {
		arg1 string
		arg2 string
	}
	displayResponseHeaderReturns struct {
		result1 error
	}
	displayResponseHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	DisplayTypeStub        func(string, time.Time) error
	displayTypeMutex       sync.RWMutex
	displayTypeArgsForCall []struct {
		arg1 string
		arg2 time.Time
	}
	displayTypeReturns struct {
		result1 error
	}
	displayTypeReturnsOnCall map[int]struct {
		result1 error
	}
	HandleInternalErrorStub        func(error)
	handleInternalErrorMutex       sync.RWMutex
	handleInternalErrorArgsForCall []struct {
		arg1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestLoggerOutput) DisplayHeader(arg1 string, arg2 string) error {
	fake.displayHeaderMutex.Lock()
	ret, specificReturn := fake.displayHeaderReturnsOnCall[len(fake.displayHeaderArgsForCall)]
	fake.displayHeaderArgsForCall = append(fake.displayHeaderArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DisplayHeader", []interface{}{arg1, arg2})
	fake.displayHeaderMutex.Unlock()
	if fake.DisplayHeaderStub != nil {
		return fake.DisplayHeaderStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayHeaderReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayHeaderCallCount() int {
	fake.displayHeaderMutex.RLock()
	defer fake.displayHeaderMutex.RUnlock()
	return len(fake.displayHeaderArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayHeaderCalls(stub func(string, string) error) {
	fake.displayHeaderMutex.Lock()
	defer fake.displayHeaderMutex.Unlock()
	fake.DisplayHeaderStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayHeaderArgsForCall(i int) (string, string) {
	fake.displayHeaderMutex.RLock()
	defer fake.displayHeaderMutex.RUnlock()
	argsForCall := fake.displayHeaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRequestLoggerOutput) DisplayHeaderReturns(result1 error) {
	fake.displayHeaderMutex.Lock()
	defer fake.displayHeaderMutex.Unlock()
	fake.DisplayHeaderStub = nil
	fake.displayHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayHeaderReturnsOnCall(i int, result1 error) {
	fake.displayHeaderMutex.Lock()
	defer fake.displayHeaderMutex.Unlock()
	fake.DisplayHeaderStub = nil
	if fake.displayHeaderReturnsOnCall == nil {
		fake.displayHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayHost(arg1 string) error {
	fake.displayHostMutex.Lock()
	ret, specificReturn := fake.displayHostReturnsOnCall[len(fake.displayHostArgsForCall)]
	fake.displayHostArgsForCall = append(fake.displayHostArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DisplayHost", []interface{}{arg1})
	fake.displayHostMutex.Unlock()
	if fake.DisplayHostStub != nil {
		return fake.DisplayHostStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayHostReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayHostCallCount() int {
	fake.displayHostMutex.RLock()
	defer fake.displayHostMutex.RUnlock()
	return len(fake.displayHostArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayHostCalls(stub func(string) error) {
	fake.displayHostMutex.Lock()
	defer fake.displayHostMutex.Unlock()
	fake.DisplayHostStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayHostArgsForCall(i int) string {
	fake.displayHostMutex.RLock()
	defer fake.displayHostMutex.RUnlock()
	argsForCall := fake.displayHostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestLoggerOutput) DisplayHostReturns(result1 error) {
	fake.displayHostMutex.Lock()
	defer fake.displayHostMutex.Unlock()
	fake.DisplayHostStub = nil
	fake.displayHostReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayHostReturnsOnCall(i int, result1 error) {
	fake.displayHostMutex.Lock()
	defer fake.displayHostMutex.Unlock()
	fake.DisplayHostStub = nil
	if fake.displayHostReturnsOnCall == nil {
		fake.displayHostReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayHostReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBody(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.displayJSONBodyMutex.Lock()
	ret, specificReturn := fake.displayJSONBodyReturnsOnCall[len(fake.displayJSONBodyArgsForCall)]
	fake.displayJSONBodyArgsForCall = append(fake.displayJSONBodyArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("DisplayJSONBody", []interface{}{arg1Copy})
	fake.displayJSONBodyMutex.Unlock()
	if fake.DisplayJSONBodyStub != nil {
		return fake.DisplayJSONBodyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayJSONBodyReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBodyCallCount() int {
	fake.displayJSONBodyMutex.RLock()
	defer fake.displayJSONBodyMutex.RUnlock()
	return len(fake.displayJSONBodyArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBodyCalls(stub func([]byte) error) {
	fake.displayJSONBodyMutex.Lock()
	defer fake.displayJSONBodyMutex.Unlock()
	fake.DisplayJSONBodyStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBodyArgsForCall(i int) []byte {
	fake.displayJSONBodyMutex.RLock()
	defer fake.displayJSONBodyMutex.RUnlock()
	argsForCall := fake.displayJSONBodyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBodyReturns(result1 error) {
	fake.displayJSONBodyMutex.Lock()
	defer fake.displayJSONBodyMutex.Unlock()
	fake.DisplayJSONBodyStub = nil
	fake.displayJSONBodyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayJSONBodyReturnsOnCall(i int, result1 error) {
	fake.displayJSONBodyMutex.Lock()
	defer fake.displayJSONBodyMutex.Unlock()
	fake.DisplayJSONBodyStub = nil
	if fake.displayJSONBodyReturnsOnCall == nil {
		fake.displayJSONBodyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayJSONBodyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayMessage(arg1 string) error {
	fake.displayMessageMutex.Lock()
	ret, specificReturn := fake.displayMessageReturnsOnCall[len(fake.displayMessageArgsForCall)]
	fake.displayMessageArgsForCall = append(fake.displayMessageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DisplayMessage", []interface{}{arg1})
	fake.displayMessageMutex.Unlock()
	if fake.DisplayMessageStub != nil {
		return fake.DisplayMessageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayMessageReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayMessageCallCount() int {
	fake.displayMessageMutex.RLock()
	defer fake.displayMessageMutex.RUnlock()
	return len(fake.displayMessageArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayMessageCalls(stub func(string) error) {
	fake.displayMessageMutex.Lock()
	defer fake.displayMessageMutex.Unlock()
	fake.DisplayMessageStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayMessageArgsForCall(i int) string {
	fake.displayMessageMutex.RLock()
	defer fake.displayMessageMutex.RUnlock()
	argsForCall := fake.displayMessageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestLoggerOutput) DisplayMessageReturns(result1 error) {
	fake.displayMessageMutex.Lock()
	defer fake.displayMessageMutex.Unlock()
	fake.DisplayMessageStub = nil
	fake.displayMessageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayMessageReturnsOnCall(i int, result1 error) {
	fake.displayMessageMutex.Lock()
	defer fake.displayMessageMutex.Unlock()
	fake.DisplayMessageStub = nil
	if fake.displayMessageReturnsOnCall == nil {
		fake.displayMessageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayMessageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeader(arg1 string, arg2 string, arg3 string) error {
	fake.displayRequestHeaderMutex.Lock()
	ret, specificReturn := fake.displayRequestHeaderReturnsOnCall[len(fake.displayRequestHeaderArgsForCall)]
	fake.displayRequestHeaderArgsForCall = append(fake.displayRequestHeaderArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DisplayRequestHeader", []interface{}{arg1, arg2, arg3})
	fake.displayRequestHeaderMutex.Unlock()
	if fake.DisplayRequestHeaderStub != nil {
		return fake.DisplayRequestHeaderStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayRequestHeaderReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeaderCallCount() int {
	fake.displayRequestHeaderMutex.RLock()
	defer fake.displayRequestHeaderMutex.RUnlock()
	return len(fake.displayRequestHeaderArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeaderCalls(stub func(string, string, string) error) {
	fake.displayRequestHeaderMutex.Lock()
	defer fake.displayRequestHeaderMutex.Unlock()
	fake.DisplayRequestHeaderStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeaderArgsForCall(i int) (string, string, string) {
	fake.displayRequestHeaderMutex.RLock()
	defer fake.displayRequestHeaderMutex.RUnlock()
	argsForCall := fake.displayRequestHeaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeaderReturns(result1 error) {
	fake.displayRequestHeaderMutex.Lock()
	defer fake.displayRequestHeaderMutex.Unlock()
	fake.DisplayRequestHeaderStub = nil
	fake.displayRequestHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayRequestHeaderReturnsOnCall(i int, result1 error) {
	fake.displayRequestHeaderMutex.Lock()
	defer fake.displayRequestHeaderMutex.Unlock()
	fake.DisplayRequestHeaderStub = nil
	if fake.displayRequestHeaderReturnsOnCall == nil {
		fake.displayRequestHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayRequestHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeader(arg1 string, arg2 string) error {
	fake.displayResponseHeaderMutex.Lock()
	ret, specificReturn := fake.displayResponseHeaderReturnsOnCall[len(fake.displayResponseHeaderArgsForCall)]
	fake.displayResponseHeaderArgsForCall = append(fake.displayResponseHeaderArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DisplayResponseHeader", []interface{}{arg1, arg2})
	fake.displayResponseHeaderMutex.Unlock()
	if fake.DisplayResponseHeaderStub != nil {
		return fake.DisplayResponseHeaderStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayResponseHeaderReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeaderCallCount() int {
	fake.displayResponseHeaderMutex.RLock()
	defer fake.displayResponseHeaderMutex.RUnlock()
	return len(fake.displayResponseHeaderArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeaderCalls(stub func(string, string) error) {
	fake.displayResponseHeaderMutex.Lock()
	defer fake.displayResponseHeaderMutex.Unlock()
	fake.DisplayResponseHeaderStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeaderArgsForCall(i int) (string, string) {
	fake.displayResponseHeaderMutex.RLock()
	defer fake.displayResponseHeaderMutex.RUnlock()
	argsForCall := fake.displayResponseHeaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeaderReturns(result1 error) {
	fake.displayResponseHeaderMutex.Lock()
	defer fake.displayResponseHeaderMutex.Unlock()
	fake.DisplayResponseHeaderStub = nil
	fake.displayResponseHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayResponseHeaderReturnsOnCall(i int, result1 error) {
	fake.displayResponseHeaderMutex.Lock()
	defer fake.displayResponseHeaderMutex.Unlock()
	fake.DisplayResponseHeaderStub = nil
	if fake.displayResponseHeaderReturnsOnCall == nil {
		fake.displayResponseHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayResponseHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayType(arg1 string, arg2 time.Time) error {
	fake.displayTypeMutex.Lock()
	ret, specificReturn := fake.displayTypeReturnsOnCall[len(fake.displayTypeArgsForCall)]
	fake.displayTypeArgsForCall = append(fake.displayTypeArgsForCall, struct {
		arg1 string
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("DisplayType", []interface{}{arg1, arg2})
	fake.displayTypeMutex.Unlock()
	if fake.DisplayTypeStub != nil {
		return fake.DisplayTypeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.displayTypeReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) DisplayTypeCallCount() int {
	fake.displayTypeMutex.RLock()
	defer fake.displayTypeMutex.RUnlock()
	return len(fake.displayTypeArgsForCall)
}

func (fake *FakeRequestLoggerOutput) DisplayTypeCalls(stub func(string, time.Time) error) {
	fake.displayTypeMutex.Lock()
	defer fake.displayTypeMutex.Unlock()
	fake.DisplayTypeStub = stub
}

func (fake *FakeRequestLoggerOutput) DisplayTypeArgsForCall(i int) (string, time.Time) {
	fake.displayTypeMutex.RLock()
	defer fake.displayTypeMutex.RUnlock()
	argsForCall := fake.displayTypeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRequestLoggerOutput) DisplayTypeReturns(result1 error) {
	fake.displayTypeMutex.Lock()
	defer fake.displayTypeMutex.Unlock()
	fake.DisplayTypeStub = nil
	fake.displayTypeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) DisplayTypeReturnsOnCall(i int, result1 error) {
	fake.displayTypeMutex.Lock()
	defer fake.displayTypeMutex.Unlock()
	fake.DisplayTypeStub = nil
	if fake.displayTypeReturnsOnCall == nil {
		fake.displayTypeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.displayTypeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) HandleInternalError(arg1 error) {
	fake.handleInternalErrorMutex.Lock()
	fake.handleInternalErrorArgsForCall = append(fake.handleInternalErrorArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.recordInvocation("HandleInternalError", []interface{}{arg1})
	fake.handleInternalErrorMutex.Unlock()
	if fake.HandleInternalErrorStub != nil {
		fake.HandleInternalErrorStub(arg1)
	}
}

func (fake *FakeRequestLoggerOutput) HandleInternalErrorCallCount() int {
	fake.handleInternalErrorMutex.RLock()
	defer fake.handleInternalErrorMutex.RUnlock()
	return len(fake.handleInternalErrorArgsForCall)
}

func (fake *FakeRequestLoggerOutput) HandleInternalErrorCalls(stub func(error)) {
	fake.handleInternalErrorMutex.Lock()
	defer fake.handleInternalErrorMutex.Unlock()
	fake.HandleInternalErrorStub = stub
}

func (fake *FakeRequestLoggerOutput) HandleInternalErrorArgsForCall(i int) error {
	fake.handleInternalErrorMutex.RLock()
	defer fake.handleInternalErrorMutex.RUnlock()
	argsForCall := fake.handleInternalErrorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestLoggerOutput) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeRequestLoggerOutput) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeRequestLoggerOutput) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopReturns
	return fakeReturns.result1
}

func (fake *FakeRequestLoggerOutput) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeRequestLoggerOutput) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeRequestLoggerOutput) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestLoggerOutput) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.displayHeaderMutex.RLock()
	defer fake.displayHeaderMutex.RUnlock()
	fake.displayHostMutex.RLock()
	defer fake.displayHostMutex.RUnlock()
	fake.displayJSONBodyMutex.RLock()
	defer fake.displayJSONBodyMutex.RUnlock()
	fake.displayMessageMutex.RLock()
	defer fake.displayMessageMutex.RUnlock()
	fake.displayRequestHeaderMutex.RLock()
	defer fake.displayRequestHeaderMutex.RUnlock()
	fake.displayResponseHeaderMutex.RLock()
	defer fake.displayResponseHeaderMutex.RUnlock()
	fake.displayTypeMutex.RLock()
	defer fake.displayTypeMutex.RUnlock()
	fake.handleInternalErrorMutex.RLock()
	defer fake.handleInternalErrorMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestLoggerOutput) recordInvocation(key string, args []interface{}) {
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

var _ wrapper.RequestLoggerOutput = new(FakeRequestLoggerOutput)
// Code generated by counterfeiter. DO NOT EDIT.
package libfakes

import (
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"gitlab.eng.vmware.com/marketplace-partner-eng/relok8s/v2/lib"

	//"gitlab.eng.vmware.com/marketplace-partner-eng/relok8s/v2/lib"

	//"gitlab.eng.vmware.com/marketplace-partner-eng/relok8s/v2/lib"
)

type FakeImageInterface struct {
	CheckStub        func(string, name.Reference) (bool, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 string
		arg2 name.Reference
	}
	checkReturns struct {
		result1 bool
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	PullStub        func(name.Reference) (v1.Image, string, error)
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		arg1 name.Reference
	}
	pullReturns struct {
		result1 v1.Image
		result2 string
		result3 error
	}
	pullReturnsOnCall map[int]struct {
		result1 v1.Image
		result2 string
		result3 error
	}
	PushStub        func(v1.Image, name.Reference) error
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		arg1 v1.Image
		arg2 name.Reference
	}
	pushReturns struct {
		result1 error
	}
	pushReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageInterface) Check(arg1 string, arg2 name.Reference) (bool, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 string
		arg2 name.Reference
	}{arg1, arg2})
	fake.recordInvocation("Check", []interface{}{arg1, arg2})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageInterface) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeImageInterface) CheckCalls(stub func(string, name.Reference) (bool, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeImageInterface) CheckArgsForCall(i int) (string, name.Reference) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageInterface) CheckReturns(result1 bool, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImageInterface) CheckReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImageInterface) Pull(arg1 name.Reference) (v1.Image, string, error) {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	fake.recordInvocation("Pull", []interface{}{arg1})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.pullReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeImageInterface) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeImageInterface) PullCalls(stub func(name.Reference) (v1.Image, string, error)) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = stub
}

func (fake *FakeImageInterface) PullArgsForCall(i int) name.Reference {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	argsForCall := fake.pullArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImageInterface) PullReturns(result1 v1.Image, result2 string, result3 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 v1.Image
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageInterface) PullReturnsOnCall(i int, result1 v1.Image, result2 string, result3 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 v1.Image
			result2 string
			result3 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 v1.Image
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageInterface) Push(arg1 v1.Image, arg2 name.Reference) error {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		arg1 v1.Image
		arg2 name.Reference
	}{arg1, arg2})
	fake.recordInvocation("Push", []interface{}{arg1, arg2})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		return fake.PushStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pushReturns
	return fakeReturns.result1
}

func (fake *FakeImageInterface) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeImageInterface) PushCalls(stub func(v1.Image, name.Reference) error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = stub
}

func (fake *FakeImageInterface) PushArgsForCall(i int) (v1.Image, name.Reference) {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	argsForCall := fake.pushArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageInterface) PushReturns(result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageInterface) PushReturnsOnCall(i int, result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageInterface) recordInvocation(key string, args []interface{}) {
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

var _ lib.ImageInterface = new(FakeImageInterface)

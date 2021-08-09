// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/IBM-Blockchain/bcdb-server/internal/comm"
	"github.com/IBM-Blockchain/bcdb-server/pkg/types"
)

type LedgerReader struct {
	GetStub        func(uint64) (*types.Block, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 uint64
	}
	getReturns struct {
		result1 *types.Block
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *types.Block
		result2 error
	}
	HeightStub        func() (uint64, error)
	heightMutex       sync.RWMutex
	heightArgsForCall []struct {
	}
	heightReturns struct {
		result1 uint64
		result2 error
	}
	heightReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LedgerReader) Get(arg1 uint64) (*types.Block, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 uint64
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LedgerReader) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *LedgerReader) GetCalls(stub func(uint64) (*types.Block, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *LedgerReader) GetArgsForCall(i int) uint64 {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LedgerReader) GetReturns(result1 *types.Block, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *types.Block
		result2 error
	}{result1, result2}
}

func (fake *LedgerReader) GetReturnsOnCall(i int, result1 *types.Block, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *types.Block
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *types.Block
		result2 error
	}{result1, result2}
}

func (fake *LedgerReader) Height() (uint64, error) {
	fake.heightMutex.Lock()
	ret, specificReturn := fake.heightReturnsOnCall[len(fake.heightArgsForCall)]
	fake.heightArgsForCall = append(fake.heightArgsForCall, struct {
	}{})
	fake.recordInvocation("Height", []interface{}{})
	fake.heightMutex.Unlock()
	if fake.HeightStub != nil {
		return fake.HeightStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.heightReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LedgerReader) HeightCallCount() int {
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	return len(fake.heightArgsForCall)
}

func (fake *LedgerReader) HeightCalls(stub func() (uint64, error)) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = stub
}

func (fake *LedgerReader) HeightReturns(result1 uint64, result2 error) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	fake.heightReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *LedgerReader) HeightReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	if fake.heightReturnsOnCall == nil {
		fake.heightReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.heightReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *LedgerReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LedgerReader) recordInvocation(key string, args []interface{}) {
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

var _ comm.LedgerReader = new(LedgerReader)

// Code generated by counterfeiter. DO NOT EDIT.
package statefakes

import (
	"io/fs"
	"os"
	"sync"

	"github.com/nginxinc/nginx-kubernetes-gateway/internal/state"
)

type FakeFileManager struct {
	ChmodStub        func(*os.File, fs.FileMode) error
	chmodMutex       sync.RWMutex
	chmodArgsForCall []struct {
		arg1 *os.File
		arg2 fs.FileMode
	}
	chmodReturns struct {
		result1 error
	}
	chmodReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(string) (*os.File, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
	}
	createReturns struct {
		result1 *os.File
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	ReadDirStub        func(string) ([]fs.DirEntry, error)
	readDirMutex       sync.RWMutex
	readDirArgsForCall []struct {
		arg1 string
	}
	readDirReturns struct {
		result1 []fs.DirEntry
		result2 error
	}
	readDirReturnsOnCall map[int]struct {
		result1 []fs.DirEntry
		result2 error
	}
	RemoveStub        func(string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func(*os.File, []byte) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 *os.File
		arg2 []byte
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileManager) Chmod(arg1 *os.File, arg2 fs.FileMode) error {
	fake.chmodMutex.Lock()
	ret, specificReturn := fake.chmodReturnsOnCall[len(fake.chmodArgsForCall)]
	fake.chmodArgsForCall = append(fake.chmodArgsForCall, struct {
		arg1 *os.File
		arg2 fs.FileMode
	}{arg1, arg2})
	stub := fake.ChmodStub
	fakeReturns := fake.chmodReturns
	fake.recordInvocation("Chmod", []interface{}{arg1, arg2})
	fake.chmodMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileManager) ChmodCallCount() int {
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	return len(fake.chmodArgsForCall)
}

func (fake *FakeFileManager) ChmodCalls(stub func(*os.File, fs.FileMode) error) {
	fake.chmodMutex.Lock()
	defer fake.chmodMutex.Unlock()
	fake.ChmodStub = stub
}

func (fake *FakeFileManager) ChmodArgsForCall(i int) (*os.File, fs.FileMode) {
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	argsForCall := fake.chmodArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFileManager) ChmodReturns(result1 error) {
	fake.chmodMutex.Lock()
	defer fake.chmodMutex.Unlock()
	fake.ChmodStub = nil
	fake.chmodReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) ChmodReturnsOnCall(i int, result1 error) {
	fake.chmodMutex.Lock()
	defer fake.chmodMutex.Unlock()
	fake.ChmodStub = nil
	if fake.chmodReturnsOnCall == nil {
		fake.chmodReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.chmodReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) Create(arg1 string) (*os.File, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFileManager) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeFileManager) CreateCalls(stub func(string) (*os.File, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeFileManager) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileManager) CreateReturns(result1 *os.File, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileManager) CreateReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileManager) ReadDir(arg1 string) ([]fs.DirEntry, error) {
	fake.readDirMutex.Lock()
	ret, specificReturn := fake.readDirReturnsOnCall[len(fake.readDirArgsForCall)]
	fake.readDirArgsForCall = append(fake.readDirArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReadDirStub
	fakeReturns := fake.readDirReturns
	fake.recordInvocation("ReadDir", []interface{}{arg1})
	fake.readDirMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFileManager) ReadDirCallCount() int {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	return len(fake.readDirArgsForCall)
}

func (fake *FakeFileManager) ReadDirCalls(stub func(string) ([]fs.DirEntry, error)) {
	fake.readDirMutex.Lock()
	defer fake.readDirMutex.Unlock()
	fake.ReadDirStub = stub
}

func (fake *FakeFileManager) ReadDirArgsForCall(i int) string {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	argsForCall := fake.readDirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileManager) ReadDirReturns(result1 []fs.DirEntry, result2 error) {
	fake.readDirMutex.Lock()
	defer fake.readDirMutex.Unlock()
	fake.ReadDirStub = nil
	fake.readDirReturns = struct {
		result1 []fs.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeFileManager) ReadDirReturnsOnCall(i int, result1 []fs.DirEntry, result2 error) {
	fake.readDirMutex.Lock()
	defer fake.readDirMutex.Unlock()
	fake.ReadDirStub = nil
	if fake.readDirReturnsOnCall == nil {
		fake.readDirReturnsOnCall = make(map[int]struct {
			result1 []fs.DirEntry
			result2 error
		})
	}
	fake.readDirReturnsOnCall[i] = struct {
		result1 []fs.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeFileManager) Remove(arg1 string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveStub
	fakeReturns := fake.removeReturns
	fake.recordInvocation("Remove", []interface{}{arg1})
	fake.removeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileManager) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeFileManager) RemoveCalls(stub func(string) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeFileManager) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileManager) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) Write(arg1 *os.File, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 *os.File
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2Copy})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileManager) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeFileManager) WriteCalls(stub func(*os.File, []byte) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeFileManager) WriteArgsForCall(i int) (*os.File, []byte) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFileManager) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFileManager) recordInvocation(key string, args []interface{}) {
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

var _ state.FileManager = new(FakeFileManager)
// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"io"
	"os"
	"sync"
	"syscall"

	"github.com/BooleanCat/go-command"
)

type Cmd struct {
	CombinedOutputStub        func() ([]byte, error)
	combinedOutputMutex       sync.RWMutex
	combinedOutputArgsForCall []struct{}
	combinedOutputReturns     struct {
		result1 []byte
		result2 error
	}
	combinedOutputReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	OutputStub        func() ([]byte, error)
	outputMutex       sync.RWMutex
	outputArgsForCall []struct{}
	outputReturns     struct {
		result1 []byte
		result2 error
	}
	outputReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RunStub        func() error
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns     struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StdinPipeStub        func() (io.WriteCloser, error)
	stdinPipeMutex       sync.RWMutex
	stdinPipeArgsForCall []struct{}
	stdinPipeReturns     struct {
		result1 io.WriteCloser
		result2 error
	}
	stdinPipeReturnsOnCall map[int]struct {
		result1 io.WriteCloser
		result2 error
	}
	StdoutPipeStub        func() (io.ReadCloser, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct{}
	stdoutPipeReturns     struct {
		result1 io.ReadCloser
		result2 error
	}
	stdoutPipeReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StderrPipeStub        func() (io.ReadCloser, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct{}
	stderrPipeReturns     struct {
		result1 io.ReadCloser
		result2 error
	}
	stderrPipeReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	GetPathStub        func() string
	getPathMutex       sync.RWMutex
	getPathArgsForCall []struct{}
	getPathReturns     struct {
		result1 string
	}
	getPathReturnsOnCall map[int]struct {
		result1 string
	}
	GetArgsStub        func() []string
	getArgsMutex       sync.RWMutex
	getArgsArgsForCall []struct{}
	getArgsReturns     struct {
		result1 []string
	}
	getArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetEnvStub        func() []string
	getEnvMutex       sync.RWMutex
	getEnvArgsForCall []struct{}
	getEnvReturns     struct {
		result1 []string
	}
	getEnvReturnsOnCall map[int]struct {
		result1 []string
	}
	GetDirStub        func() string
	getDirMutex       sync.RWMutex
	getDirArgsForCall []struct{}
	getDirReturns     struct {
		result1 string
	}
	getDirReturnsOnCall map[int]struct {
		result1 string
	}
	GetStdinStub        func() io.Reader
	getStdinMutex       sync.RWMutex
	getStdinArgsForCall []struct{}
	getStdinReturns     struct {
		result1 io.Reader
	}
	getStdinReturnsOnCall map[int]struct {
		result1 io.Reader
	}
	GetStdoutStub        func() io.Writer
	getStdoutMutex       sync.RWMutex
	getStdoutArgsForCall []struct{}
	getStdoutReturns     struct {
		result1 io.Writer
	}
	getStdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	GetStderrStub        func() io.Writer
	getStderrMutex       sync.RWMutex
	getStderrArgsForCall []struct{}
	getStderrReturns     struct {
		result1 io.Writer
	}
	getStderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	GetExtraFilesStub        func() []*os.File
	getExtraFilesMutex       sync.RWMutex
	getExtraFilesArgsForCall []struct{}
	getExtraFilesReturns     struct {
		result1 []*os.File
	}
	getExtraFilesReturnsOnCall map[int]struct {
		result1 []*os.File
	}
	GetSysProcAttrStub        func() *syscall.SysProcAttr
	getSysProcAttrMutex       sync.RWMutex
	getSysProcAttrArgsForCall []struct{}
	getSysProcAttrReturns     struct {
		result1 *syscall.SysProcAttr
	}
	getSysProcAttrReturnsOnCall map[int]struct {
		result1 *syscall.SysProcAttr
	}
	GetProcessStub        func() command.Process
	getProcessMutex       sync.RWMutex
	getProcessArgsForCall []struct{}
	getProcessReturns     struct {
		result1 command.Process
	}
	getProcessReturnsOnCall map[int]struct {
		result1 command.Process
	}
	GetProcessStateStub        func() *os.ProcessState
	getProcessStateMutex       sync.RWMutex
	getProcessStateArgsForCall []struct{}
	getProcessStateReturns     struct {
		result1 *os.ProcessState
	}
	getProcessStateReturnsOnCall map[int]struct {
		result1 *os.ProcessState
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Cmd) CombinedOutput() ([]byte, error) {
	fake.combinedOutputMutex.Lock()
	ret, specificReturn := fake.combinedOutputReturnsOnCall[len(fake.combinedOutputArgsForCall)]
	fake.combinedOutputArgsForCall = append(fake.combinedOutputArgsForCall, struct{}{})
	fake.recordInvocation("CombinedOutput", []interface{}{})
	fake.combinedOutputMutex.Unlock()
	if fake.CombinedOutputStub != nil {
		return fake.CombinedOutputStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.combinedOutputReturns.result1, fake.combinedOutputReturns.result2
}

func (fake *Cmd) CombinedOutputCallCount() int {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return len(fake.combinedOutputArgsForCall)
}

func (fake *Cmd) CombinedOutputReturns(result1 []byte, result2 error) {
	fake.CombinedOutputStub = nil
	fake.combinedOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Cmd) CombinedOutputReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.CombinedOutputStub = nil
	if fake.combinedOutputReturnsOnCall == nil {
		fake.combinedOutputReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.combinedOutputReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Cmd) Output() ([]byte, error) {
	fake.outputMutex.Lock()
	ret, specificReturn := fake.outputReturnsOnCall[len(fake.outputArgsForCall)]
	fake.outputArgsForCall = append(fake.outputArgsForCall, struct{}{})
	fake.recordInvocation("Output", []interface{}{})
	fake.outputMutex.Unlock()
	if fake.OutputStub != nil {
		return fake.OutputStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.outputReturns.result1, fake.outputReturns.result2
}

func (fake *Cmd) OutputCallCount() int {
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	return len(fake.outputArgsForCall)
}

func (fake *Cmd) OutputReturns(result1 []byte, result2 error) {
	fake.OutputStub = nil
	fake.outputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Cmd) OutputReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.OutputStub = nil
	if fake.outputReturnsOnCall == nil {
		fake.outputReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.outputReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Cmd) Run() error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *Cmd) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *Cmd) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *Cmd) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Cmd) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *Cmd) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *Cmd) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *Cmd) StartReturnsOnCall(i int, result1 error) {
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

func (fake *Cmd) StdinPipe() (io.WriteCloser, error) {
	fake.stdinPipeMutex.Lock()
	ret, specificReturn := fake.stdinPipeReturnsOnCall[len(fake.stdinPipeArgsForCall)]
	fake.stdinPipeArgsForCall = append(fake.stdinPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdinPipe", []interface{}{})
	fake.stdinPipeMutex.Unlock()
	if fake.StdinPipeStub != nil {
		return fake.StdinPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stdinPipeReturns.result1, fake.stdinPipeReturns.result2
}

func (fake *Cmd) StdinPipeCallCount() int {
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	return len(fake.stdinPipeArgsForCall)
}

func (fake *Cmd) StdinPipeReturns(result1 io.WriteCloser, result2 error) {
	fake.StdinPipeStub = nil
	fake.stdinPipeReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) StdinPipeReturnsOnCall(i int, result1 io.WriteCloser, result2 error) {
	fake.StdinPipeStub = nil
	if fake.stdinPipeReturnsOnCall == nil {
		fake.stdinPipeReturnsOnCall = make(map[int]struct {
			result1 io.WriteCloser
			result2 error
		})
	}
	fake.stdinPipeReturnsOnCall[i] = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) StdoutPipe() (io.ReadCloser, error) {
	fake.stdoutPipeMutex.Lock()
	ret, specificReturn := fake.stdoutPipeReturnsOnCall[len(fake.stdoutPipeArgsForCall)]
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdoutPipe", []interface{}{})
	fake.stdoutPipeMutex.Unlock()
	if fake.StdoutPipeStub != nil {
		return fake.StdoutPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stdoutPipeReturns.result1, fake.stdoutPipeReturns.result2
}

func (fake *Cmd) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

func (fake *Cmd) StdoutPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) StdoutPipeReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.StdoutPipeStub = nil
	if fake.stdoutPipeReturnsOnCall == nil {
		fake.stdoutPipeReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.stdoutPipeReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) StderrPipe() (io.ReadCloser, error) {
	fake.stderrPipeMutex.Lock()
	ret, specificReturn := fake.stderrPipeReturnsOnCall[len(fake.stderrPipeArgsForCall)]
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct{}{})
	fake.recordInvocation("StderrPipe", []interface{}{})
	fake.stderrPipeMutex.Unlock()
	if fake.StderrPipeStub != nil {
		return fake.StderrPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stderrPipeReturns.result1, fake.stderrPipeReturns.result2
}

func (fake *Cmd) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

func (fake *Cmd) StderrPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) StderrPipeReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.StderrPipeStub = nil
	if fake.stderrPipeReturnsOnCall == nil {
		fake.stderrPipeReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.stderrPipeReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *Cmd) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitReturns.result1
}

func (fake *Cmd) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *Cmd) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *Cmd) WaitReturnsOnCall(i int, result1 error) {
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Cmd) GetPath() string {
	fake.getPathMutex.Lock()
	ret, specificReturn := fake.getPathReturnsOnCall[len(fake.getPathArgsForCall)]
	fake.getPathArgsForCall = append(fake.getPathArgsForCall, struct{}{})
	fake.recordInvocation("GetPath", []interface{}{})
	fake.getPathMutex.Unlock()
	if fake.GetPathStub != nil {
		return fake.GetPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getPathReturns.result1
}

func (fake *Cmd) GetPathCallCount() int {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	return len(fake.getPathArgsForCall)
}

func (fake *Cmd) GetPathReturns(result1 string) {
	fake.GetPathStub = nil
	fake.getPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *Cmd) GetPathReturnsOnCall(i int, result1 string) {
	fake.GetPathStub = nil
	if fake.getPathReturnsOnCall == nil {
		fake.getPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Cmd) GetArgs() []string {
	fake.getArgsMutex.Lock()
	ret, specificReturn := fake.getArgsReturnsOnCall[len(fake.getArgsArgsForCall)]
	fake.getArgsArgsForCall = append(fake.getArgsArgsForCall, struct{}{})
	fake.recordInvocation("GetArgs", []interface{}{})
	fake.getArgsMutex.Unlock()
	if fake.GetArgsStub != nil {
		return fake.GetArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getArgsReturns.result1
}

func (fake *Cmd) GetArgsCallCount() int {
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	return len(fake.getArgsArgsForCall)
}

func (fake *Cmd) GetArgsReturns(result1 []string) {
	fake.GetArgsStub = nil
	fake.getArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *Cmd) GetArgsReturnsOnCall(i int, result1 []string) {
	fake.GetArgsStub = nil
	if fake.getArgsReturnsOnCall == nil {
		fake.getArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *Cmd) GetEnv() []string {
	fake.getEnvMutex.Lock()
	ret, specificReturn := fake.getEnvReturnsOnCall[len(fake.getEnvArgsForCall)]
	fake.getEnvArgsForCall = append(fake.getEnvArgsForCall, struct{}{})
	fake.recordInvocation("GetEnv", []interface{}{})
	fake.getEnvMutex.Unlock()
	if fake.GetEnvStub != nil {
		return fake.GetEnvStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getEnvReturns.result1
}

func (fake *Cmd) GetEnvCallCount() int {
	fake.getEnvMutex.RLock()
	defer fake.getEnvMutex.RUnlock()
	return len(fake.getEnvArgsForCall)
}

func (fake *Cmd) GetEnvReturns(result1 []string) {
	fake.GetEnvStub = nil
	fake.getEnvReturns = struct {
		result1 []string
	}{result1}
}

func (fake *Cmd) GetEnvReturnsOnCall(i int, result1 []string) {
	fake.GetEnvStub = nil
	if fake.getEnvReturnsOnCall == nil {
		fake.getEnvReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getEnvReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *Cmd) GetDir() string {
	fake.getDirMutex.Lock()
	ret, specificReturn := fake.getDirReturnsOnCall[len(fake.getDirArgsForCall)]
	fake.getDirArgsForCall = append(fake.getDirArgsForCall, struct{}{})
	fake.recordInvocation("GetDir", []interface{}{})
	fake.getDirMutex.Unlock()
	if fake.GetDirStub != nil {
		return fake.GetDirStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getDirReturns.result1
}

func (fake *Cmd) GetDirCallCount() int {
	fake.getDirMutex.RLock()
	defer fake.getDirMutex.RUnlock()
	return len(fake.getDirArgsForCall)
}

func (fake *Cmd) GetDirReturns(result1 string) {
	fake.GetDirStub = nil
	fake.getDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *Cmd) GetDirReturnsOnCall(i int, result1 string) {
	fake.GetDirStub = nil
	if fake.getDirReturnsOnCall == nil {
		fake.getDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Cmd) GetStdin() io.Reader {
	fake.getStdinMutex.Lock()
	ret, specificReturn := fake.getStdinReturnsOnCall[len(fake.getStdinArgsForCall)]
	fake.getStdinArgsForCall = append(fake.getStdinArgsForCall, struct{}{})
	fake.recordInvocation("GetStdin", []interface{}{})
	fake.getStdinMutex.Unlock()
	if fake.GetStdinStub != nil {
		return fake.GetStdinStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStdinReturns.result1
}

func (fake *Cmd) GetStdinCallCount() int {
	fake.getStdinMutex.RLock()
	defer fake.getStdinMutex.RUnlock()
	return len(fake.getStdinArgsForCall)
}

func (fake *Cmd) GetStdinReturns(result1 io.Reader) {
	fake.GetStdinStub = nil
	fake.getStdinReturns = struct {
		result1 io.Reader
	}{result1}
}

func (fake *Cmd) GetStdinReturnsOnCall(i int, result1 io.Reader) {
	fake.GetStdinStub = nil
	if fake.getStdinReturnsOnCall == nil {
		fake.getStdinReturnsOnCall = make(map[int]struct {
			result1 io.Reader
		})
	}
	fake.getStdinReturnsOnCall[i] = struct {
		result1 io.Reader
	}{result1}
}

func (fake *Cmd) GetStdout() io.Writer {
	fake.getStdoutMutex.Lock()
	ret, specificReturn := fake.getStdoutReturnsOnCall[len(fake.getStdoutArgsForCall)]
	fake.getStdoutArgsForCall = append(fake.getStdoutArgsForCall, struct{}{})
	fake.recordInvocation("GetStdout", []interface{}{})
	fake.getStdoutMutex.Unlock()
	if fake.GetStdoutStub != nil {
		return fake.GetStdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStdoutReturns.result1
}

func (fake *Cmd) GetStdoutCallCount() int {
	fake.getStdoutMutex.RLock()
	defer fake.getStdoutMutex.RUnlock()
	return len(fake.getStdoutArgsForCall)
}

func (fake *Cmd) GetStdoutReturns(result1 io.Writer) {
	fake.GetStdoutStub = nil
	fake.getStdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *Cmd) GetStdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.GetStdoutStub = nil
	if fake.getStdoutReturnsOnCall == nil {
		fake.getStdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.getStdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *Cmd) GetStderr() io.Writer {
	fake.getStderrMutex.Lock()
	ret, specificReturn := fake.getStderrReturnsOnCall[len(fake.getStderrArgsForCall)]
	fake.getStderrArgsForCall = append(fake.getStderrArgsForCall, struct{}{})
	fake.recordInvocation("GetStderr", []interface{}{})
	fake.getStderrMutex.Unlock()
	if fake.GetStderrStub != nil {
		return fake.GetStderrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStderrReturns.result1
}

func (fake *Cmd) GetStderrCallCount() int {
	fake.getStderrMutex.RLock()
	defer fake.getStderrMutex.RUnlock()
	return len(fake.getStderrArgsForCall)
}

func (fake *Cmd) GetStderrReturns(result1 io.Writer) {
	fake.GetStderrStub = nil
	fake.getStderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *Cmd) GetStderrReturnsOnCall(i int, result1 io.Writer) {
	fake.GetStderrStub = nil
	if fake.getStderrReturnsOnCall == nil {
		fake.getStderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.getStderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *Cmd) GetExtraFiles() []*os.File {
	fake.getExtraFilesMutex.Lock()
	ret, specificReturn := fake.getExtraFilesReturnsOnCall[len(fake.getExtraFilesArgsForCall)]
	fake.getExtraFilesArgsForCall = append(fake.getExtraFilesArgsForCall, struct{}{})
	fake.recordInvocation("GetExtraFiles", []interface{}{})
	fake.getExtraFilesMutex.Unlock()
	if fake.GetExtraFilesStub != nil {
		return fake.GetExtraFilesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getExtraFilesReturns.result1
}

func (fake *Cmd) GetExtraFilesCallCount() int {
	fake.getExtraFilesMutex.RLock()
	defer fake.getExtraFilesMutex.RUnlock()
	return len(fake.getExtraFilesArgsForCall)
}

func (fake *Cmd) GetExtraFilesReturns(result1 []*os.File) {
	fake.GetExtraFilesStub = nil
	fake.getExtraFilesReturns = struct {
		result1 []*os.File
	}{result1}
}

func (fake *Cmd) GetExtraFilesReturnsOnCall(i int, result1 []*os.File) {
	fake.GetExtraFilesStub = nil
	if fake.getExtraFilesReturnsOnCall == nil {
		fake.getExtraFilesReturnsOnCall = make(map[int]struct {
			result1 []*os.File
		})
	}
	fake.getExtraFilesReturnsOnCall[i] = struct {
		result1 []*os.File
	}{result1}
}

func (fake *Cmd) GetSysProcAttr() *syscall.SysProcAttr {
	fake.getSysProcAttrMutex.Lock()
	ret, specificReturn := fake.getSysProcAttrReturnsOnCall[len(fake.getSysProcAttrArgsForCall)]
	fake.getSysProcAttrArgsForCall = append(fake.getSysProcAttrArgsForCall, struct{}{})
	fake.recordInvocation("GetSysProcAttr", []interface{}{})
	fake.getSysProcAttrMutex.Unlock()
	if fake.GetSysProcAttrStub != nil {
		return fake.GetSysProcAttrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getSysProcAttrReturns.result1
}

func (fake *Cmd) GetSysProcAttrCallCount() int {
	fake.getSysProcAttrMutex.RLock()
	defer fake.getSysProcAttrMutex.RUnlock()
	return len(fake.getSysProcAttrArgsForCall)
}

func (fake *Cmd) GetSysProcAttrReturns(result1 *syscall.SysProcAttr) {
	fake.GetSysProcAttrStub = nil
	fake.getSysProcAttrReturns = struct {
		result1 *syscall.SysProcAttr
	}{result1}
}

func (fake *Cmd) GetSysProcAttrReturnsOnCall(i int, result1 *syscall.SysProcAttr) {
	fake.GetSysProcAttrStub = nil
	if fake.getSysProcAttrReturnsOnCall == nil {
		fake.getSysProcAttrReturnsOnCall = make(map[int]struct {
			result1 *syscall.SysProcAttr
		})
	}
	fake.getSysProcAttrReturnsOnCall[i] = struct {
		result1 *syscall.SysProcAttr
	}{result1}
}

func (fake *Cmd) GetProcess() command.Process {
	fake.getProcessMutex.Lock()
	ret, specificReturn := fake.getProcessReturnsOnCall[len(fake.getProcessArgsForCall)]
	fake.getProcessArgsForCall = append(fake.getProcessArgsForCall, struct{}{})
	fake.recordInvocation("GetProcess", []interface{}{})
	fake.getProcessMutex.Unlock()
	if fake.GetProcessStub != nil {
		return fake.GetProcessStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getProcessReturns.result1
}

func (fake *Cmd) GetProcessCallCount() int {
	fake.getProcessMutex.RLock()
	defer fake.getProcessMutex.RUnlock()
	return len(fake.getProcessArgsForCall)
}

func (fake *Cmd) GetProcessReturns(result1 command.Process) {
	fake.GetProcessStub = nil
	fake.getProcessReturns = struct {
		result1 command.Process
	}{result1}
}

func (fake *Cmd) GetProcessReturnsOnCall(i int, result1 command.Process) {
	fake.GetProcessStub = nil
	if fake.getProcessReturnsOnCall == nil {
		fake.getProcessReturnsOnCall = make(map[int]struct {
			result1 command.Process
		})
	}
	fake.getProcessReturnsOnCall[i] = struct {
		result1 command.Process
	}{result1}
}

func (fake *Cmd) GetProcessState() *os.ProcessState {
	fake.getProcessStateMutex.Lock()
	ret, specificReturn := fake.getProcessStateReturnsOnCall[len(fake.getProcessStateArgsForCall)]
	fake.getProcessStateArgsForCall = append(fake.getProcessStateArgsForCall, struct{}{})
	fake.recordInvocation("GetProcessState", []interface{}{})
	fake.getProcessStateMutex.Unlock()
	if fake.GetProcessStateStub != nil {
		return fake.GetProcessStateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getProcessStateReturns.result1
}

func (fake *Cmd) GetProcessStateCallCount() int {
	fake.getProcessStateMutex.RLock()
	defer fake.getProcessStateMutex.RUnlock()
	return len(fake.getProcessStateArgsForCall)
}

func (fake *Cmd) GetProcessStateReturns(result1 *os.ProcessState) {
	fake.GetProcessStateStub = nil
	fake.getProcessStateReturns = struct {
		result1 *os.ProcessState
	}{result1}
}

func (fake *Cmd) GetProcessStateReturnsOnCall(i int, result1 *os.ProcessState) {
	fake.GetProcessStateStub = nil
	if fake.getProcessStateReturnsOnCall == nil {
		fake.getProcessStateReturnsOnCall = make(map[int]struct {
			result1 *os.ProcessState
		})
	}
	fake.getProcessStateReturnsOnCall[i] = struct {
		result1 *os.ProcessState
	}{result1}
}

func (fake *Cmd) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	fake.getEnvMutex.RLock()
	defer fake.getEnvMutex.RUnlock()
	fake.getDirMutex.RLock()
	defer fake.getDirMutex.RUnlock()
	fake.getStdinMutex.RLock()
	defer fake.getStdinMutex.RUnlock()
	fake.getStdoutMutex.RLock()
	defer fake.getStdoutMutex.RUnlock()
	fake.getStderrMutex.RLock()
	defer fake.getStderrMutex.RUnlock()
	fake.getExtraFilesMutex.RLock()
	defer fake.getExtraFilesMutex.RUnlock()
	fake.getSysProcAttrMutex.RLock()
	defer fake.getSysProcAttrMutex.RUnlock()
	fake.getProcessMutex.RLock()
	defer fake.getProcessMutex.RUnlock()
	fake.getProcessStateMutex.RLock()
	defer fake.getProcessStateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Cmd) recordInvocation(key string, args []interface{}) {
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

var _ command.Cmd = new(Cmd)

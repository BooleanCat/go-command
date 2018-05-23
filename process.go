package command

import "os"

type BaseProcess interface {
	Kill() error
	Release() error
	Signal(os.Signal) error
	Wait() (*os.ProcessState, error)
}

//go:generate counterfeiter -o commandfakes/process.go -fake-name Process . Process
type Process interface {
	BaseProcess

	GetPid() int
	SetPid(int) Process
}

type ProcessShim struct {
	process *os.Process
}

func NewProcess(process *os.Process) *ProcessShim {
	return &ProcessShim{process: process}
}

func (p *ProcessShim) Kill() error {
	return p.process.Kill()
}

func (p *ProcessShim) Release() error {
	return p.process.Release()
}

func (p *ProcessShim) Signal(sig os.Signal) error {
	return p.process.Signal(sig)
}

func (p *ProcessShim) Wait() (*os.ProcessState, error) {
	return p.process.Wait()
}

func (p *ProcessShim) GetPid() int {
	return p.process.Pid
}

func (p *ProcessShim) SetPid(pid int) Process {
	p.process.Pid = pid
	return p
}

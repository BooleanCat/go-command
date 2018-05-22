package command

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

type BaseCmd interface {
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	Run() error
	Start() error
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Wait() error
}

//go:generate counterfeiter -o commandfakes/cmd.go . Cmd
type Cmd interface {
	BaseCmd

	GetPath() string
	GetArgs() []string
	GetEnv() []string
	GetDir() string
	GetStdin() io.Reader
	GetStdout() io.Writer
	GetStderr() io.Writer
	GetExtraFiles() []*os.File
	GetSysProcAttr() *syscall.SysProcAttr
	GetProcess() Process
	GetProcessState() *os.ProcessState
}

type Shim struct {
	cmd *exec.Cmd
}

func Wrap(c *exec.Cmd) Cmd {
	return &Shim{cmd: c}
}

func (s *Shim) CombinedOutput() ([]byte, error) {
	return s.cmd.CombinedOutput()
}

func (s *Shim) Output() ([]byte, error) {
	return s.cmd.Output()
}

func (s *Shim) Run() error {
	return s.cmd.Run()
}

func (s *Shim) Start() error {
	return s.cmd.Start()
}

func (s *Shim) StdinPipe() (io.WriteCloser, error) {
	return s.cmd.StdinPipe()
}

func (s *Shim) StdoutPipe() (io.ReadCloser, error) {
	return s.cmd.StdoutPipe()
}

func (s *Shim) StderrPipe() (io.ReadCloser, error) {
	return s.cmd.StderrPipe()
}

func (s *Shim) Wait() error {
	return s.cmd.Wait()
}

func (s *Shim) GetPath() string {
	return s.cmd.Path
}

func (s *Shim) GetArgs() []string {
	return s.cmd.Args
}

func (s *Shim) GetEnv() []string {
	return s.cmd.Env
}

func (s *Shim) GetDir() string {
	return s.cmd.Dir
}

func (s *Shim) GetStdin() io.Reader {
	return s.cmd.Stdin
}

func (s *Shim) GetStdout() io.Writer {
	return s.cmd.Stdout
}

func (s *Shim) GetStderr() io.Writer {
	return s.cmd.Stderr
}

func (s *Shim) GetExtraFiles() []*os.File {
	return s.cmd.ExtraFiles
}

func (s *Shim) GetSysProcAttr() *syscall.SysProcAttr {
	return s.cmd.SysProcAttr
}

func (s *Shim) GetProcess() Process {
	return NewProcess(s.cmd.Process)
}
func (s *Shim) GetProcessState() *os.ProcessState {
	return s.cmd.ProcessState
}

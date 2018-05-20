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

//go:generate counterfeiter . Cmd
type Cmd interface {
	BaseCmd

	GetPath() string
	SetPath(string) Cmd

	GetArgs() []string
	SetArgs([]string) Cmd

	GetEnv() []string
	SetEnv([]string) Cmd

	GetDir() string
	SetDir(string) Cmd

	GetStdin() io.Reader
	SetStdin(io.Reader) Cmd

	GetStdout() io.Writer
	SetStdout(io.Writer) Cmd

	GetStderr() io.Writer
	SetStderr(io.Writer) Cmd

	GetExtraFiles() []*os.File
	SetExtraFiles([]*os.File) Cmd

	GetSysProcAttr() *syscall.SysProcAttr
	SetSysProcAttr(*syscall.SysProcAttr) Cmd

	GetProcess() *os.Process
	SetProcess(process *os.Process) Cmd

	GetProcessState() *os.ProcessState
	SetProcessState(*os.ProcessState) Cmd
}

type Shim struct {
	cmd *exec.Cmd
}

func Command(path string, arg ...string) Cmd {
	return &Shim{cmd: exec.Command(path, arg...)}
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

func (s *Shim) SetPath(path string) Cmd {
	s.cmd.Path = path
	return s
}

func (s *Shim) GetArgs() []string {
	return s.cmd.Args
}

func (s *Shim) SetArgs(args []string) Cmd {
	s.cmd.Args = args
	return s
}

func (s *Shim) GetEnv() []string {
	return s.cmd.Env
}

func (s *Shim) SetEnv(env []string) Cmd {
	s.cmd.Env = env
	return s
}

func (s *Shim) GetDir() string {
	return s.cmd.Dir
}

func (s *Shim) SetDir(dir string) Cmd {
	s.cmd.Dir = dir
	return s
}

func (s *Shim) GetStdin() io.Reader {
	return s.cmd.Stdin
}

func (s *Shim) SetStdin(stdin io.Reader) Cmd {
	s.cmd.Stdin = stdin
	return s
}

func (s *Shim) GetStdout() io.Writer {
	return s.cmd.Stdout
}

func (s *Shim) SetStdout(stdout io.Writer) Cmd {
	s.cmd.Stdout = stdout
	return s
}

func (s *Shim) GetStderr() io.Writer {
	return s.cmd.Stderr
}

func (s *Shim) SetStderr(stderr io.Writer) Cmd {
	s.cmd.Stderr = stderr
	return s
}

func (s *Shim) GetExtraFiles() []*os.File {
	return s.cmd.ExtraFiles
}

func (s *Shim) SetExtraFiles(extraFiles []*os.File) Cmd {
	s.cmd.ExtraFiles = extraFiles
	return s
}

func (s *Shim) GetSysProcAttr() *syscall.SysProcAttr {
	return s.cmd.SysProcAttr
}

func (s *Shim) SetSysProcAttr(sysProcAttr *syscall.SysProcAttr) Cmd {
	s.cmd.SysProcAttr = sysProcAttr
	return s
}

func (s *Shim) GetProcess() *os.Process {
	return s.cmd.Process
}

func (s *Shim) SetProcess(process *os.Process) Cmd {
	s.cmd.Process = process
	return s
}

func (s *Shim) GetProcessState() *os.ProcessState {
	return s.cmd.ProcessState
}

func (s *Shim) SetProcessState(processState *os.ProcessState) Cmd {
	s.cmd.ProcessState = processState
	return s
}

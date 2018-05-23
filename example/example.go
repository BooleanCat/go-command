package example

import (
	"os/exec"

	command "github.com/BooleanCat/go-command"
	"golang.org/x/sys/unix"
)

type Sleeper struct {
	cmd         command.Cmd
	commandWrap command.Wrapper
}

func NewSleeper() *Sleeper {
	return &Sleeper{commandWrap: command.Wrap}
}

func (s *Sleeper) Start() error {
	s.cmd = s.commandWrap(exec.Command("sleep", "3600"))
	return s.cmd.Start()
}

func (s *Sleeper) Stop() error {
	return s.cmd.GetProcess().Signal(unix.SIGTERM)
}

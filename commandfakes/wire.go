package commandfakes

func Wired() (*Wrap, *Cmd, *Process) {
	wrap := new(Wrap)
	cmd := new(Cmd)
	process := new(Process)

	wrap.Returns(cmd)
	cmd.GetProcessReturns(process)

	return wrap, cmd, process
}

package command_test

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"

	command "github.com/BooleanCat/go-command"
	fakes "github.com/BooleanCat/go-command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoCommand", func() {
	Describe("BaseCmd", func() {
		It("is implemented by exec.Cmd", func() {
			_ = command.BaseCmd(new(exec.Cmd))
		})
	})

	Describe("Cmd", func() {
		It("is implemented by Shim", func() {
			_ = command.Cmd(new(command.Shim))
		})

		It("is implemented by FakeCmd", func() {
			_ = command.Cmd(new(fakes.FakeCmd))
		})
	})

	Describe("Shim", func() {
		var (
			cmd    command.Cmd
			buffer *bytes.Buffer
		)

		BeforeEach(func() {
			cmd = command.Command("foo")
			buffer = new(bytes.Buffer)
		})

		Describe("SetPath/GetPath", func() {
			It("delegates to the underlying cmd.Path", func() {
				cmd = cmd.SetPath("bar")
				Expect(cmd.GetPath()).To(Equal("bar"))
			})
		})

		Describe("SetArgs/GetArgs", func() {
			It("delegates to the underlying cmd.Args", func() {
				args := []string{"bar", "baz"}
				cmd = cmd.SetArgs(args)
				Expect(cmd.GetArgs()).To(Equal(args))
			})
		})

		Describe("SetEnv/GetEnv", func() {
			It("delegates to the underlying cmd.Env", func() {
				env := []string{"foo=bar"}
				cmd = cmd.SetEnv(env)
				Expect(cmd.GetEnv()).To(Equal(env))
			})
		})

		Describe("SetDir/GetDir", func() {
			It("delegates to the underlying cmd.Dir", func() {
				cmd = cmd.SetDir("/foo/bar")
				Expect(cmd.GetDir()).To(Equal("/foo/bar"))
			})
		})

		Describe("SetStdin/GetStdin", func() {
			It("delegates to the underlying cmd.Stdin", func() {
				stdin := cmd.SetStdin(buffer).GetStdin()
				writeString(buffer, "foo")
				Expect(string(readAll(stdin))).To(Equal("foo"))
			})
		})

		Describe("SetStdout/GetStdout", func() {
			It("delegates to the underlying cmd.Stdout", func() {
				stdout := cmd.SetStdout(buffer).GetStdout()
				writeString(stdout, "foo")
				Expect(string(readAll(buffer))).To(Equal("foo"))
			})
		})

		Describe("SetStderr/GetStderr", func() {
			It("delegates to the underlying cmd.Stderr", func() {
				stderr := cmd.SetStderr(buffer).GetStderr()
				writeString(stderr, "foo")
				Expect(string(readAll(buffer))).To(Equal("foo"))
			})
		})

		Describe("SetExtraFiles/GetExtraFiles", func() {
			It("delegates to the underlying cmd.ExtraFiles", func() {
				extraFiles := []*os.File{new(os.File)}
				cmd.SetExtraFiles(extraFiles)
				Expect(cmd.GetExtraFiles()).To(HaveLen(1))
				Expect(cmd.GetExtraFiles()[0]).To(BeIdenticalTo(extraFiles[0]))
			})
		})

		Describe("SetSysProcAttr/GetSysProcAttr", func() {
			It("delegates to the underlying cmd.ExtraSysProcAttr", func() {
				sysProcAttr := new(syscall.SysProcAttr)
				cmd.SetSysProcAttr(sysProcAttr)
				Expect(cmd.GetSysProcAttr()).To(BeIdenticalTo(sysProcAttr))
			})
		})

		Describe("SetProcess/GetProcess", func() {
			It("delegates to the underlying cmd.Process", func() {
				process := &os.Process{Pid: 4}
				cmd.SetProcess(process)
				Expect(cmd.GetProcess().GetPid()).To(Equal(4))
			})
		})

		Describe("SetProcessState/GetProcessState", func() {
			It("delegates to the underlying cmd.ProcessState", func() {
				processState := new(os.ProcessState)
				cmd = cmd.SetProcessState(processState)
				Expect(cmd.GetProcessState()).To(BeIdenticalTo(processState))
			})
		})
	})
})

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

	Describe("Wrapper", func() {
		It("is implemented by Wrap", func() {
			_ = command.Wrapper(command.Wrap)
		})

		It("is implemented by fakes.Wrap", func() {
			_ = command.Wrapper(new(fakes.Wrap).Spy)
		})
	})

	Describe("Cmd", func() {
		It("is implemented by Shim", func() {
			_ = command.Cmd(new(command.Shim))
		})

		It("is implemented by commandfakes.Cmd", func() {
			_ = command.Cmd(new(fakes.Cmd))
		})
	})

	Describe("Shim", func() {
		var (
			execCmd *exec.Cmd
			cmd     command.Cmd
			buffer  *bytes.Buffer
		)

		BeforeEach(func() {
			buffer = new(bytes.Buffer)
			execCmd = exec.Command("foo", "bar")
		})

		JustBeforeEach(func() {
			cmd = command.Wrap(execCmd)
		})

		Describe("GetPath", func() {
			It("delegates to the underlying cmd.Path", func() {
				Expect(cmd.GetPath()).To(Equal("foo"))
			})
		})

		Describe("GetArgs", func() {
			It("delegates to the underlying cmd.Args", func() {
				Expect(cmd.GetArgs()).To(Equal([]string{"foo", "bar"}))
			})
		})

		Describe("GetEnv", func() {
			BeforeEach(func() {
				execCmd.Env = []string{"foo=bar"}
			})

			It("delegates to the underlying cmd.Env", func() {
				Expect(cmd.GetEnv()).To(Equal([]string{"foo=bar"}))
			})
		})

		Describe("GetDir", func() {
			BeforeEach(func() {
				execCmd.Dir = "/foo/bar"
			})

			It("delegates to the underlying cmd.Dir", func() {
				Expect(cmd.GetDir()).To(Equal("/foo/bar"))
			})
		})

		Describe("GetStdin", func() {
			BeforeEach(func() {
				execCmd.Stdin = buffer
			})

			It("delegates to the underlying cmd.Stdin", func() {
				writeString(buffer, "foo")
				stdin := readAll(cmd.GetStdin())
				Expect(string(stdin)).To(Equal("foo"))
			})
		})

		Describe("GetStdout", func() {
			BeforeEach(func() {
				execCmd.Stdout = buffer
			})

			It("delegates to the underlying cmd.Stdout", func() {
				writeString(cmd.GetStdout(), "foo")
				Expect(buffer.String()).To(Equal("foo"))
			})
		})

		Describe("GetStderr", func() {
			BeforeEach(func() {
				execCmd.Stderr = buffer
			})

			It("delegates to the underlying cmd.Stderr", func() {
				writeString(cmd.GetStderr(), "foo")
				Expect(buffer.String()).To(Equal("foo"))
			})
		})

		Describe("GetExtraFiles", func() {
			var extraFiles []*os.File

			BeforeEach(func() {
				extraFiles = []*os.File{new(os.File)}
				execCmd.ExtraFiles = extraFiles
			})

			It("delegates to the underlying cmd.ExtraFiles", func() {
				Expect(cmd.GetExtraFiles()).To(HaveLen(1))
				Expect(cmd.GetExtraFiles()[0]).To(BeIdenticalTo(extraFiles[0]))
			})
		})

		Describe("GetSysProcAttr", func() {
			var sysProcAttr *syscall.SysProcAttr

			BeforeEach(func() {
				sysProcAttr = new(syscall.SysProcAttr)
				execCmd.SysProcAttr = sysProcAttr
			})

			It("delegates to the underlying cmd.ExtraSysProcAttr", func() {
				Expect(cmd.GetSysProcAttr()).To(BeIdenticalTo(sysProcAttr))
			})
		})

		Describe("GetProcess", func() {
			var process *os.Process

			BeforeEach(func() {
				process = &os.Process{Pid: 42}
				execCmd.Process = process
			})

			It("delegates to the underlying cmd.Process", func() {
				Expect(cmd.GetProcess().GetPid()).To(Equal(42))
			})
		})

		Describe("GetProcessState", func() {
			var processState *os.ProcessState

			BeforeEach(func() {
				processState = new(os.ProcessState)
				execCmd.ProcessState = processState
			})

			It("delegates to the underlying cmd.ProcessState", func() {
				Expect(cmd.GetProcessState()).To(BeIdenticalTo(processState))
			})
		})
	})
})

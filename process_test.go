package command_test

import (
	"os"

	command "github.com/BooleanCat/go-command"
	fakes "github.com/BooleanCat/go-command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoCommand", func() {
	Describe("BaseProcess", func() {
		It("is implemented by os.Process", func() {
			_ = command.BaseProcess(new(os.Process))
		})
	})

	Describe("Process", func() {
		It("is implemented by ProcessShim", func() {
			_ = command.Process(new(command.ProcessShim))
		})

		It("is implemented by commandfakes.Process", func() {
			_ = command.Process(new(fakes.Process))
		})
	})

	Describe("ProcessShim", func() {
		Describe("SetPid/GetPid", func() {
			It("delegates to the underlying process.Pid", func() {
				process := command.NewProcess(new(os.Process)).SetPid(42)
				Expect(process.GetPid()).To(Equal(42))
			})
		})
	})
})

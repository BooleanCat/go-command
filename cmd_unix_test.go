package command_test

import (
	"os/exec"

	command "github.com/BooleanCat/go-command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoCommand", func() {
	Describe("Shim", func() {
		It("delegates to the underlying exec.Cmd", func() {
			cmd := command.Wrap(exec.Command("echo", "hello"))
			output, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(output)).To(Equal("hello\n"))
		})
	})
})

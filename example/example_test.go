package example

import (
	fakes "github.com/BooleanCat/go-command/commandfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/sys/unix"
)

var _ = Describe("Sleeper", func() {
	var (
		sleeper     *Sleeper
		fakeWrap    *fakes.Wrap
		fakeCmd     *fakes.Cmd
		fakeProcess *fakes.Process
	)

	BeforeEach(func() {
		fakeWrap, fakeCmd, fakeProcess = fakes.Wired()
		sleeper = NewSleeper()
	})

	AfterEach(func() {
		sleeper.Stop()
	})

	Describe("Start", func() {
		It("succeeds", func() {
			Expect(sleeper.Start()).NotTo(HaveOccurred())
		})

		It("executes a sleep command", func() {
			sleeper.commandWrap = fakeWrap.Spy
			Expect(sleeper.Start()).To(Succeed())
			Expect(fakeCmd.StartCallCount()).To(Equal(1))
		})
	})

	Describe("Stop", func() {
		It("succeeds", func() {
			Expect(sleeper.Start()).To(Succeed())
			Expect(sleeper.Stop()).To(Succeed())
		})

		It("sends SIGTERM to the sleep process", func() {
			sleeper.commandWrap = fakeWrap.Spy
			Expect(sleeper.Start()).To(Succeed())
			Expect(sleeper.Stop()).To(Succeed())
			Expect(fakeProcess.SignalCallCount()).To(Equal(1))
			Expect(fakeProcess.SignalArgsForCall(0)).To(Equal(unix.SIGTERM))
		})
	})
})

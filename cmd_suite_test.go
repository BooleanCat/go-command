package command_test

import (
	"io"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCommand Suite")
}

func readAll(r io.Reader) []byte {
	content, err := ioutil.ReadAll(r)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return content
}

func writeString(w io.Writer, s string) {
	n, err := io.WriteString(w, s)
	ExpectWithOffset(1, n).To(Equal(len(s)))
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
}

package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Macadam starts", Label("windows", "linux", "darwin", "noVM"), func() {

	It("non-existing VM", func() {
		startCMD := []string{"start", "123"}
		code, _, err := runCMD(startCMD)
		Expect(code).ShouldNot(Equal(0))
		Expect(err).Should(Equal("VM \"123\" does not exist"))
	})
})

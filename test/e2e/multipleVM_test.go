package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Macadam init setup test", Label("multiple", "linux", "darwin"), func() {
	BeforeEach(func() {
		vmNumberCheck(0)
	})

	AfterEach(func() {
		removeAllVM()
	})

	It("create multiple CentOS VM", Label("mul-centos"), func() {
		// init a CentOS VM with name vm1
		initCMD := []string{"init", "--name", "vm1", IMAGE}
		runCMDsuccess(initCMD)

		// check the list command returns one item
		vmNumberCheck(1)

		// start the CentOS VM1
		startCMD := []string{"start", "vm1"}
		output := runCMDsuccess(startCMD)
		Expect(output).Should(ContainSubstring("started successfully"))

		// ssh into the VM1 and prints user
		sshCMD := []string{"ssh", "vm1", "whoami"}
		output = runCMDsuccess(sshCMD)
		Expect(output).Should(Equal("core"))

		//init another centos VM with name vm2
		initCMD2 := []string{"init", "--name", "vm2", IMAGE}
		runCMDsuccess(initCMD2)

		// check the list command returns two items
		vmNumberCheck(2)

		// start the CentOS VM2
		startCMD2 := []string{"start", "vm2"}
		output = runCMDsuccess(startCMD2)
		Expect(output).Should(ContainSubstring("started successfully"))

		// ssh into the VM2 and prints user
		sshCMD2 := []string{"ssh", "vm2", "whoami"}
		output = runCMDsuccess(sshCMD2)
		Expect(output).Should(Equal("core"))

		//check again VM1 still running
		output = runCMDsuccess(sshCMD)
		Expect(output).Should(Equal("core"))
	})

})

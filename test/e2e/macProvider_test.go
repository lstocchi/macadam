package e2e

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("provider libkrun test with centos", Label("libkrun", "darwin"), func() {
	BeforeEach(func() {
		vmNumberCheck(0, "libkrun")
	})

	AfterEach(func() {
		removeAllVM("libkrun")
	})

	It("init CentOS VM with cpu, disk and memory setup", Label("hardware"), func() {
		testparam := Init_Hardware_Parameter{
			cpu:           "3",
			disk:          "30",
			memory:        "2048",
			expect_memory: []string{"1.7G", "1.8G", "1.9G", "2G"},
		}
		init_hardware_test(IMAGE, testparam, "libkrun")
	})

	It("init CentOS VM with username and sshkey setup", Label("sshkey"), func() {
		init_sshkey_test(IMAGE, "libkrun")
	})

	It("init CentOS VM with cloud-init setup", Label("cloudinit"), func() {
		init_cloudinit_test(IMAGE, "libkrun")
	})

	It("init CentOS VM with name", Label("name"), func() {
		init_vmName_test(IMAGE, "libkrun")
	})

})

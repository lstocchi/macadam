package e2e

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("provider hyperv test with vhd image", Label("hyperv", "windows"), func() {
	BeforeEach(func() {
		if !strings.HasSuffix(IMAGE, ".vhd") && !strings.HasSuffix(IMAGE, ".vhdx") {
			Skip("Skipping Hyper-V tests because image is not VHD/VHDX")
		}
		vmNumberCheck(0, "hyperv")
	})

	AfterEach(func() {
		removeAllVM("hyperv")
	})

	It("init VM with cpu, disk and memory setup", Label("hardware"), func() {
		testparam := Init_Hardware_Parameter{
			cpu:           "3",
			disk:          "30",
			memory:        "2048",
			expect_memory: []string{"1.7G", "1.8G", "1.9G", "2G"},
		}
		init_hardware_test(IMAGE, testparam, "hyperv")
	})

	It("init VM with username and sshkey setup", Label("sshkey"), func() {
		init_sshkey_test(IMAGE, "hyperv")
	})

	It("init VM with cloud-init setup", Label("cloudinit"), func() {
		init_cloudinit_test(IMAGE, "hyperv")
	})

	It("init VM with name", Label("name"), func() {
		init_vmName_test(IMAGE, "hyperv")
	})

	It("init VM without parameter", Label("noparameter"), func() {
		init_vm_no_param_test(IMAGE, "hyperv")
	})

})

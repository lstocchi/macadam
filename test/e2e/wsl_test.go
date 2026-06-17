package e2e

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("provider wsl test with wsl image", Label("wsl", "windows"), func() {
	BeforeEach(func() {
		if !strings.HasSuffix(IMAGE, ".wsl") && !strings.HasSuffix(IMAGE, ".tar.gz") {
			Skip("Skipping WSL tests because image is not wsl or tar.gz")
		}
		vmNumberCheck(0, "wsl")
	})

	AfterEach(func() {
		removeAllVM("wsl")
	})

	It("init VM with cpu, disk and memory setup", Label("hardware"), func() {
		testparam := Init_Hardware_Parameter{
			cpu:           "3",
			disk:          "30",
			memory:        "2048",
			expect_memory: []string{"1.7G", "1.8G", "1.9G", "2G"},
		}
		init_hardware_test(IMAGE, testparam, "wsl")
	})

	It("init VM with username and sshkey setup", Label("sshkey"), func() {
		init_sshkey_test(IMAGE, "wsl")
	})

	It("init VM with cloud-init setup", Label("cloudinit"), func() {
		init_cloudinit_test(IMAGE, "wsl")
	})

	It("init VM with name", Label("name"), func() {
		init_vmName_test(IMAGE, "wsl")
	})

	It("init VM without parameter", Label("noparameter"), func() {
		init_vm_no_param_test(IMAGE, "wsl")
	})
})

package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Macadam init setup test", Label("init", "linux", "darwin"), func() {
	BeforeEach(func() {
		vmNumberCheck(0)
	})

	AfterEach(func() {
		removeAllVM()
	})

	tests := Init_Test_Para{
		{
			cpu:           "1",
			disk:          "20",
			memory:        "3072",
			expect_memory: []string{"2.6G", "2.7G", "2.7G", "2.9G"},
		},
		{
			cpu:           "3",
			disk:          "30",
			memory:        "2048",
			expect_memory: []string{"1.7G", "1.8G", "1.9G", "2G"},
		},
		{
			cpu:           "4",
			disk:          "10",
			memory:        "4096",
			expect_memory: []string{"3.5G", "3.6G", "3.7G", "3.8G"},
		},
	}
	for _, tt := range tests {
		It(fmt.Sprintf("init CentOS VM with cpu=%s, disk=%s, memory=%s", tt.cpu, tt.disk, tt.memory), Label("hardware"), func() {
			init_hardware_test(IMAGE, tt)
		})
	}

	It("init CentOS VM with username and sshkey setup", Label("sshkey"), func() {
		init_sshkey_test(IMAGE)
	})

	It("init CentOS VM with cloud-init setup", Label("cloudinit"), func() {
		// init a CentOS VM with cpu and disk-size setup
		init_cloudinit_test(IMAGE)
	})

	It("init CentOS VM with name", Label("name"), func() {
		// init a CentOS VM with name setup
		init_vmName_test(IMAGE)
	})
})

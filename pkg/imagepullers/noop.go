package imagepullers

import (
	"github.com/containers/podman/v5/pkg/machine/define"
)

type NoopImagePuller struct {
	localPath string
}

var _ define.ImagePuller = &NoopImagePuller{}

func (puller *NoopImagePuller) SetSourceURI(localPath string) {
	puller.localPath = localPath
}

func (puller *NoopImagePuller) LocalPath() (*define.VMFile, error) {
	return define.NewMachineFile(puller.localPath, nil)
}

func (puller *NoopImagePuller) Download() error {
	return nil
}

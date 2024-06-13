package main

import (
	"fmt"
	"log/slog"

	"github.com/crc-org/macadam/pkg/cmdline"

	_ "github.com/containers/podman/v5/pkg/machine"
	_ "github.com/spf13/cobra"
)

func main() {
	slog.Info(fmt.Sprintf("macadam version %s", cmdline.Version()))
}

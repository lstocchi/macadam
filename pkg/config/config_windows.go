//go:build windows

package config

import "os"

const (
	// _configPath is the path to the macadam/machines.conf
	// inside a given config directory.
	_configPath = "\\macadam\\machines.conf"
)

// userConfigPath returns the path to the users local config that is
// not shared with other users. It uses $APPDATA/containers...
func UserConfigPath() (string, error) {
	return os.Getenv("APPDATA") + _configPath, nil
}

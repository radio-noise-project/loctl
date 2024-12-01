package init

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigurationDirPath() string {
	var osName string = runtime.GOOS
	var path string
	if osName == "linux" {
		xdg := os.Getenv("XDG_CONFIG_HOME")
		if xdg == "" {
			path = filepath.Join(os.Getenv("HOME"), ".config", "loctl")
		} else {
			path = filepath.Join(xdg, "loctl")
		}
	}
	return path
}

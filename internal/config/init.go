package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func InitializeConfiguration(host string, port int) error {
	osName := runtime.GOOS
	err := getConfigurationFilePath(osName, host, port)
	return err
}

func getConfigurationFilePath(osName string, host string, port int) error {
	if osName == "linux" {
		xdg := os.Getenv("XDG_CONFIG_HOME")

		fileInfo, err := os.Lstat("./")
		if err != nil {
			return err
		}
		fileMode := fileInfo.Mode()
		unixPerms := fileMode & os.ModePerm

		if xdg == "" {
			slog.Info("XDG_CONFIG_HOME is not found")
			home := os.Getenv("HOME")
			xdg = filepath.Join(home, ".config")

			if !isDir(xdg) {
				err = os.Mkdir((xdg), unixPerms)
				if err != nil {
					return err
				}
			}
		}
		configDirPath := filepath.Join(xdg, "loctl")
		err = os.Mkdir(configDirPath, unixPerms)
		if err != nil {
			return err
		}
		configFilePath := filepath.Join(configDirPath, "config.toml")
		err = ConfigEncode(host, port, configFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

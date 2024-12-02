package version

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/BurntSushi/toml"
)

func GetVersionLoctl() (string, string, string) {
	type config struct {
		CodeName  string `toml:"codeName"`
		Version   string `toml:"version"`
		BuiltDate string `toml:"builtDate"`
	}

	c := map[string]config{}

	if _, err := toml.DecodeFile("../../VERSION.toml", &c); err != nil {
		panic(err)
	}

	return c["version"].CodeName, c["version"].Version, c["version"].BuiltDate
}

func getGolangVersion() string {
	version := runtime.Version()
	return version
}

func getOsArchVersion() (string, string) {
	os := runtime.GOOS
	arch := runtime.GOARCH
	return os, arch
}

func getGitCommitHash() string {
	var hash string
	info, err := debug.ReadBuildInfo()
	if !err {
		fmt.Print("Nothing build information")
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			hash = s.Value
		}
	}
	return hash
}

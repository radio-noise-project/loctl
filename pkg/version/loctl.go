package version

import (
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

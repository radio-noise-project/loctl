package config

type ConfigureSetting struct {
	LastOrderIpAddress string `toml:"host"`
	LastOrderPort      int    `toml:"port"`
}

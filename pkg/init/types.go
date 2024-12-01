package init

type ConfigureSetting struct {
	LastOrderIpAddress string `toml:"host"`
	LastOrderPort      int    `toml:"port"`
}

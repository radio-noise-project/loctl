package client

import (
	"net"
	"strconv"

	"github.com/radio-noise-project/loctl/pkg/config"
)

type Config struct {
	DestinationBaseUrl string
	DestinationPort    int
	DestinationUrl     string
	Timeout            int
	Token              string
}

func NewConfiguration() *Config {
	config, err := config.ConfigDecode(config.GetConfigurationFilePath())
	if err != nil {
		panic(err)
	}
	host := config["lastOrder"].LastOrderIpAddress
	port := config["lastOrder"].LastOrderPort
	dstUrl := net.JoinHostPort(host, strconv.Itoa(port))
	return &Config{
		DestinationBaseUrl: host,
		DestinationPort:    port,
		DestinationUrl:     "http://" + dstUrl,
		Timeout:            10000,
	}
}

func (c *Config) SetAuthToken(token string) *Config {
	c.Token = token
	return c
}

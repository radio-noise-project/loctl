package client

import (
	"net"
	"strconv"
)

type Config struct {
	DestinationBaseUrl string
	DestinationPort    int
	DestinationUrl     string
	Timeout            int
	Token              string
}

func NewConfiguration(host string, port int, timeout int) *Config {
	dstUrl := net.JoinHostPort(host, strconv.Itoa(port))
	return &Config{
		DestinationBaseUrl: host,
		DestinationPort:    port,
		DestinationUrl:     dstUrl,
		Timeout:            timeout,
	}
}

func (c *Config) SetAuthToken(token string) *Config {
	c.Token = token
	return c
}

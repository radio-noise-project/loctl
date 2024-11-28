package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func Initialize(config Config) *resty.Client {
	client := resty.New()
	client.SetHeader("Accept", "application/json")
	client.SetBaseURL(config.DestinationBaseUrl)
	client.SetTimeout(time.Duration(config.Timeout))
	if config.Token != "" {
		client.SetAuthToken(config.Token)
	}

	return client
}

func Get(client resty.Client, uri string, params map[string]string) (*resty.Response, error) {
	resp, err := client.R().SetQueryParams(params).Get(uri)
	return resp, err
}

func Post(client resty.Client, uri string, params map[string]string) (*resty.Response, error) {
	resp, err := client.R().SetQueryParams(params).Post(uri)
	return resp, err
}

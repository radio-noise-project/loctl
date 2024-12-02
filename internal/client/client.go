package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func (c Config) NewClient() *resty.Client {
	client := resty.New()
	client.SetHeader("Accept", "application/json")
	client.SetBaseURL(c.DestinationUrl)
	client.SetTimeout(time.Duration(c.Timeout))
	if c.Token != "" {
		client.SetAuthToken(c.Token)
	}

	return client
}

func Get(client *resty.Client, uri string, params map[string]string) (*resty.Response, error) {
	resp, err := client.R().SetQueryParams(params).Get(uri)
	return resp, err
}

func Post(client *resty.Client, uri string, params map[string]string, body interface{}) (*resty.Response, error) {
	resp, err := client.R().SetQueryParams(params).SetBody(body).Post(uri)
	return resp, err
}

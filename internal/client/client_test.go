package client

import "testing"

func TestInitialize(t *testing.T) {
	var config Config
	config.Initialize("localhost", 10032, 30)
	client := Initialize(config)
	if client.BaseURL != config.DestinationUrl{
		t.Errorf("excepted %s; but got %s", config.DestinationBaseUrl, client.BaseURL)	
	}
}

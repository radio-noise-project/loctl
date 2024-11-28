package client

import "testing"

func TestInitialize(t *testing.T) {
	config := NewConfiguration("localhost", 10032, 30)
	client := config.NewClient()
	if client.BaseURL != config.DestinationUrl {
		t.Errorf("excepted %s; but got %s", config.DestinationBaseUrl, client.BaseURL)
	}
}

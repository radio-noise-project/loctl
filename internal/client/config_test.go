package client

import (
	"strconv"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	host := "localhost"
	port := 10023
	timeout := 30
	config := NewConfiguration(host, port, timeout)
	url := host + ":" + strconv.Itoa(port)

	if config.DestinationUrl != url {
		t.Fatalf("excepted %s; but got %s", url, config.DestinationUrl)
	}

	if config.Timeout != timeout {
		t.Fatalf("excepted %d; but got %d", timeout, config.Timeout)
	}
}

func TestSetAuthToken(t *testing.T) {
	config := NewConfiguration("localhost", 10023, 30)
	token := "testToken"
	config.SetAuthToken(token)
	if config.Token != token {
		t.Fatalf("excepted %s; but got %s", token, config.Token)
	}
}

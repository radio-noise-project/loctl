package config

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func TestConfigEncode(t *testing.T) {
	address := "localhost"
	port := 8080
	path := "testdata/config.toml"
	err := ConfigEncode(address, port, path)
	if err != nil {
		t.Fatal(err)
	}

	//var testConf ConfigureSetting
	testConf := map[string]ConfigureSetting{}
	_, err = toml.DecodeFile(path, &testConf)
	if err != nil {
		t.Fatal(err)
	}

	if testConf["LastOrder"].LastOrderIpAddress != address {
		t.Errorf("excepted address = %s; but got %s", address, testConf["LastOrder"].LastOrderIpAddress)
	}
	if testConf["LastOrder"].LastOrderPort != port {
		t.Errorf("excepted port = %d; but got %d", port, testConf["LastOrder"].LastOrderPort)
	}
}

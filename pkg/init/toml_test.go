package init

import (
	"testing"
)

func TestMain(t *testing.T) {
	address := "localhost"
	port := 8080
	path := "testdata/config.toml"
	err := ConfigEncode(address, port, path)
	if err != nil {
		t.Fatalf("excepted no error; but got %v", err)
	}
	testConf, err := ConfigDecode(path)
	if err != nil {
		t.Fatalf("excepted no error; but got %v", err)
	}

	if testConf["lastOrder"].LastOrderIpAddress != address {
		t.Errorf("excepted address = %s; but got %s", address, testConf["lastOrder"].LastOrderIpAddress)
	}
	if testConf["lastOrder"].LastOrderPort != port {
		t.Errorf("excepted port = %d; but got %d", port, testConf["lastOrder"].LastOrderPort)
	}
}

func TestConfigEncode(t *testing.T) {
	address := "localhost"
	port := 8080
	path := "testdata/config.toml"

	err := ConfigEncode(address, port, path)
	if err != nil {
		t.Fatalf("excepted no error; but got %v", err)
	}
}

func TestConfigDecode(t *testing.T) {
	path := "testdata/config.toml"
	_, err := ConfigDecode(path)
	if err != nil {
		t.Fatalf("excepted no error; but god %v", err)
	}
}

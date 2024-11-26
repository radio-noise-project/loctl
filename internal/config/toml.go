package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

func ConfigEncode(address string, port int, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	config := ConfigureSetting{
		LastOrderIpAddress: address,
		LastOrderPort:      port,
	}
	configMap := map[string]ConfigureSetting{"LastOrder": config}

	err = toml.NewEncoder(file).Encode(configMap)
	if err != nil {
		return err
	}

	return nil
}

func ConfigDecode(filePath string) (map[string]ConfigureSetting, error) {
	conf := map[string]ConfigureSetting{}
	_, err := toml.DecodeFile(filePath, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

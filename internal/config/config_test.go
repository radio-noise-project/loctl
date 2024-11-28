package config

import (
	"fmt"
	"testing"
)

func TestGetConfigurationDirPath(t *testing.T) {
	path := GetConfigurationDirPath()
	fmt.Printf("Configuration path is %s", path)
}

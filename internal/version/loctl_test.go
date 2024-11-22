package version

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShowVersionLoctl(t *testing.T) {
	apath, _ := filepath.Abs("../")
	os.Chdir(apath)
	ShowVersionLoctl()
}

package version

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/radio-noise-project/loctl/internal/client"
)

func ShowVersionLastOrder() {
	info := VersionInformationLastOrder.getVersionInformation(VersionInformationLastOrder{}, "lo", 8080)

	fmt.Println("Server: Last-Order")
	fmt.Printf("CodeName: %s\n", info.CodeName)
	fmt.Printf("Version: %s\n", info.Version)
	fmt.Printf("Go version: %s\n", info.GolangVersion)
	fmt.Printf("Git commit: %s\n", info.BuiltGitCommitHash)
	fmt.Printf("Built: %s\n", info.BuiltDate)
	fmt.Printf("OS: %s\n", info.Os)
	fmt.Printf("Arch: %s\n", info.Arch)
}

func (VersionInformationLastOrder) getVersionInformation(url string, port int) *VersionInformationLastOrder {
	var versionInfo client.ResponseLastOrderVersion

	conf := client.HttpConfig.HttpConfigConstructor(client.HttpConfig{}, url, port)
	resp := client.GetVersion(*conf, "Last-Order")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &versionInfo); err != nil {
		panic(err)
	}

	info := &VersionInformationLastOrder{
		CodeName:           versionInfo.CodeName,
		Version:            versionInfo.Version,
		BuiltDate:          versionInfo.BuiltDate,
		GolangVersion:      versionInfo.GolangVersion,
		BuiltGitCommitHash: versionInfo.BuiltGitCommitHash,
		Os:                 versionInfo.Os,
		Arch:               versionInfo.Arch,
	}
	return info
}

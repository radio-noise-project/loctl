package version

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ShowVersionLoctl() {
	info := VersionInformationLoctl.getVersionInformation(VersionInformationLoctl{})

	fmt.Println("Control: loctl")
	fmt.Printf("CodeName: %s\n", info.CodeName)
	fmt.Printf("Version: %s\n", info.Version)
	fmt.Printf("Go version: %s\n", info.GolangVersion)
	fmt.Printf("Git commit: %s\n", info.BuiltGitCommitHash)
	fmt.Printf("Built: %s\n", info.BuiltDate)
	fmt.Printf("OS: %s\n", info.Os)
	fmt.Printf("Arch: %s\n", info.Arch)
}

func (VersionInformationLoctl) getVersionInformation() *VersionInformationLoctl {
	codeName, version, builtDate := getLoctlVersion()
	golangVerson := getGolangVersion()
	revisionHash := getGitCommitHash()
	os, arch := getOsArchVersion()
	info := &VersionInformationLoctl{
		CodeName:           codeName,
		Version:            version,
		BuiltDate:          strToTime(builtDate),
		GolangVersion:      golangVerson,
		BuiltGitCommitHash: revisionHash,
		Os:                 os,
		Arch:               arch,
	}
	return info
}

func getLoctlVersion() (string, string, string) {
	fp, err := os.Open("../VERSION")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var str [3]string
	var i = 0
	for scanner.Scan() {
		str[i] = scanner.Text()
		i++
	}

	codeName := str[0]
	version := str[1]
	builtDate := str[2]
	return codeName, version, builtDate
}

func strToTime(t string) time.Time {
	prasedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", t)
	if err != nil {
		panic(err)
	}

	return prasedTime
}

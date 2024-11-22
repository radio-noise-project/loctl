package client

import "time"

type ResponseLastOrderVersion struct {
	CodeName            string
	Version             string
	GolangVersion       string
	DockerEngineVersion string
	BuiltGitCommitHash  string
	BuiltDate           time.Time
	Os                  string
	Arch                string
}

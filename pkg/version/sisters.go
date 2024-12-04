package version

import (
	"encoding/json"

	"github.com/radio-noise-project/loctl/internal/client"
	"github.com/radio-noise-project/loctl/pkg/version/types"
)

func GetVersionSisters(id string, config *client.Config) (string, string, string) {
	cli := config.NewClient()
	params := map[string]string{
		"sisterId": id,
	}
	jsonResp, err := client.Get(cli, "/api/v0/runtime/version/sisters", params)
	if err != nil {
		panic(err)
	}

	var resp *types.ResponseSistersVersion
	err = json.Unmarshal(jsonResp.Body(), &resp)
	if err != nil {
		panic(err)
	}

	return resp.Version, resp.CodeName, resp.Version
}

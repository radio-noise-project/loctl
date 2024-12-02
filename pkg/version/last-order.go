package version

import (
	"encoding/json"

	"github.com/radio-noise-project/loctl/internal/client"
	"github.com/radio-noise-project/loctl/pkg/version/types"
)

func GetVersionLastOrder(config *client.Config) (string, string, string) {
	cli := config.NewClient()
	jsonResp, err := client.Get(cli, "/api/v0/runtime/version/last-order", nil)
	if err != nil {
		panic(err)
	}

	var resp types.ResponseLastOrderVersion
	err = json.Unmarshal(jsonResp.Body(), &resp)
	if err != nil {
		panic(err)
	}

	return resp.Version, resp.CodeName, resp.Version
}

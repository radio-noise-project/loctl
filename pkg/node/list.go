package node

import (
	"encoding/json"

	"github.com/radio-noise-project/loctl/internal/client"
	"github.com/radio-noise-project/loctl/pkg/node/types"
)

func ListNode() (string, string, string) {
	config := client.NewConfiguration()
	cli := config.NewClient()

	jsonResp, err := client.Get(cli, "/api/v0/node/list", nil)
	if err != nil {
		panic(err)
	}

	var resp types.ResponseNodeList
	err = json.Unmarshal(jsonResp.Body(), &resp)
	if err != nil {
		panic(err)
	}

	return resp.Name, resp.Address, resp.Port
}

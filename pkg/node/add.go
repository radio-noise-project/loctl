package node

import (
	"encoding/json"

	"github.com/radio-noise-project/loctl/internal/client"
)

func AddNode(name string, address string, port string) int {
	config := client.NewConfiguration()
	cli := config.NewClient()
	body := map[string]string{
		"name":    name,
		"address": address,
		"port":    port,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	resp, err := client.Post(cli, "/api/v0/node/add", nil, string(jsonBody))
	if err != nil {
		panic(err)
	}

	return resp.StatusCode()
}

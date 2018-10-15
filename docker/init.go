package docker

import (
	"os"

	"github.com/docker/docker/client"
)

var Client *client.Client

func init() {
	cli, err := client.NewEnvClient()
	if err != nil {
		os.Exit(0)
	}
	Client = cli
}

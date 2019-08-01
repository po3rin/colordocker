package colordocker

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

var Client *client.Client

func init() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("missing to create Docker API client")
		os.Exit(0)
	}
	Client = cli
}

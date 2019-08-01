package main

import (
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func main() {
	plugin.Run(func(dockerCli command.Cli) *cobra.Command {
		// TODO: impliments
	}, manager.Metadata{
		SchemaVersion: "0.0.1",
		Vendor:        "po3rin",
		Version:       "0.0.1",
		Experimental:  true,
	})
}

func newRootCmd(use string, dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Short:       "Color Docker",
		Long:        `A tool to color docker command.`,
		Use:         use,
		Annotations: map[string]string{"experimentalCLI": "true"},
	}
	addCommands(cmd, dockerCli)
	return cmd
}

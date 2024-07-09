package cmd

import "github.com/ignite/cli/v28/ignite/services/plugin"

// GetCommands returns the list of gui app commands.
func GetCommands() []*plugin.Command {
	return []*plugin.Command{
		{
			Use:   "gui [command]",
			Short: "Integration of the ignite-ui desktop app",
			Commands: []*plugin.Command{
				{
					Use:   "install",
					Short: "Install ignite-ui.",
					Flags: []*plugin.Flag{
						{Name: "downloadIn", Type: plugin.FlagTypeString, Usage: "directory where the ingite-ui will be downloaded"},
					},
				},
				{
					Use:   "run",
					Short: "Run ignite-ui.",
				},
			},
		},
	}
}

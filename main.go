package main

import (
	"context"
	"fmt"

	hplugin "github.com/hashicorp/go-plugin"

	"gui/cmd"

	"github.com/ignite/cli/v28/ignite/services/plugin"
)

type app struct{}

func (app) Manifest(_ context.Context) (*plugin.Manifest, error) {
	return &plugin.Manifest{
		Name:     "gui",
		Commands: cmd.GetCommands(),
	}, nil
}

func (app) Execute(ctx context.Context, c *plugin.ExecutedCommand, _ plugin.ClientAPI) error {
	// Remove the first two elements "ignite" and "gui" from OsArgs.
	args := c.OsArgs[2:]

	switch args[0] {
	case "install":
		return cmd.InstallGui(ctx, c)
	case "run":
		return cmd.RunGui(ctx, c)
	default:
		return fmt.Errorf("unknown command: %s", c.Path)
	}
}

func (app) ExecuteHookPre(_ context.Context, _ *plugin.ExecutedHook, _ plugin.ClientAPI) error {
	return nil
}

func (app) ExecuteHookPost(_ context.Context, _ *plugin.ExecutedHook, _ plugin.ClientAPI) error {
	return nil
}

func (app) ExecuteHookCleanUp(_ context.Context, _ *plugin.ExecutedHook, _ plugin.ClientAPI) error {
	return nil
}

func main() {
	hplugin.Serve(&hplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig(),
		Plugins: map[string]hplugin.Plugin{
			"gui": plugin.NewGRPC(&app{}),
		},
		GRPCServer: hplugin.DefaultGRPCServer,
	})
}

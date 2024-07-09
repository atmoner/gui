package cmd

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/ignite/cli/v28/ignite/services/plugin"
)

func RunGui(ctx context.Context, cmd *plugin.ExecutedCommand) error {
	fmt.Println("RunGui")
	cmdRun := exec.Command("./ignite-ui-0.1.6.AppImage", "--no-sandbox")
	err := cmdRun.Run()
	if err != nil {
		fmt.Printf("run error: %v\n", err)
	}
	return nil
}

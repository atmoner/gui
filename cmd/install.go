package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ignite/cli/v28/ignite/services/plugin"
)

func InstallGui(ctx context.Context, cmd *plugin.ExecutedCommand) error {
	fmt.Println("Start download app!")
	fullURLFile := "https://github.com/atmoner/ignite-ui/releases/download/v0.1.6/ignite-ui-0.1.6.AppImage"

	// Build fileName from fullPath
	fileURL, _ := url.Parse(fullURLFile)
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d \n", fileName, size)

	// Make file executable
	cmdRun := exec.Command("chmod", "+x", "ignite-ui-0.1.6.AppImage")
	cmdRun.Run()

	fmt.Printf("Ignite ui is installed and ready to use! \n")
	fmt.Printf("Run: ignite gui run \n")
	return nil
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

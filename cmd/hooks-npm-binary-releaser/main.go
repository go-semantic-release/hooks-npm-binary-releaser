package main

import (
	"log"
	"os"

	hooksReleaser "github.com/go-semantic-release/hooks-npm-binary-releaser/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
)

var VERSION = "dev"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Hooks: func() hooks.Hooks {
			return &hooksReleaser.NpmBinaryReleaser{
				PluginVersion: VERSION,
				Logger:        log.New(os.Stdout, "[npm-binary-releaser]: ", 0),
			}
		},
	})
}

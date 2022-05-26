package hooks

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/christophwitzko/npm-binary-releaser/pkg/config"
	"github.com/christophwitzko/npm-binary-releaser/pkg/releaser"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"gopkg.in/yaml.v3"
)

type NpmBinaryReleaser struct {
	PluginVersion string
	Logger        *log.Logger
}

func (t *NpmBinaryReleaser) Init(_ map[string]string) error {
	return nil
}

func (t *NpmBinaryReleaser) Name() string {
	return "npm-binary-releaser"
}

func (t *NpmBinaryReleaser) Version() string {
	return t.PluginVersion
}

func (t *NpmBinaryReleaser) Success(releaseConfig *hooks.SuccessHookConfig) error {
	var cfg config.Config
	if configData, err := os.ReadFile(".npm-binary-releaser.yaml"); err == nil {
		t.Logger.Println("reading config from .npm-binary-releaser.yaml")
		if err := yaml.Unmarshal(configData, &cfg); err != nil {
			return fmt.Errorf("could not read config: %w", err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	envInfo := config.GetRepositoryAndHomepageFromEnv()
	if cfg.Repository == "" {
		cfg.Repository = envInfo.Repository
	}
	if cfg.Homepage == "" {
		cfg.Homepage = envInfo.Homepage
	}
	if cfg.BinName == "" {
		cfg.BinName = envInfo.PackageName
	}
	if cfg.InputBinDirPath == "" {
		cfg.TryDefaultInputPaths = true
	}
	if cfg.OutputDirPath == "" {
		cfg.OutputDirPath = config.DefaultOutputDirPath
	}
	if cfg.PublishRegistry == "" {
		cfg.PublishRegistry = config.DefaultPublishRegistry
	}
	cfg.PackageVersion = releaseConfig.NewRelease.Version

	return releaser.Run(&cfg, t.Logger)
}

func (t *NpmBinaryReleaser) NoRelease(_ *hooks.NoReleaseConfig) error {
	return nil
}

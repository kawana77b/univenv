package config

import (
	"path/filepath"

	"github.com/kawana77b/univenv/internal/config/item"
	"github.com/kawana77b/univenv/internal/sysutil"
	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
	"gopkg.in/yaml.v3"
)

const (
	ENV_UNIVENV_CONFIG_DIR string = "UNIVENV_CONFIG_DIR"
	DEFAULT_FILE_BASENAME  string = "config"
	DEFAULT_FILE_EXT       string = ".yml"
)

var DEFAULT_DIR = filepath.Join(sysutil.HomeDir(), ".config", "univenv")

type Config struct {
	Items []item.Item `yaml:"items"`
}

func (c *Config) GetEnableItems(sh shell.Shell, os ostype.OS) []item.Item {
	return item.GetEnabledItems(c.Items, os, sh)
}

func (c *Config) Read(b []byte) error {
	err := yaml.Unmarshal(b, c)
	if err != nil {
		return err
	}
	return c.validate()
}

func (c *Config) Write() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	return yaml.Marshal(c)
}

func (c *Config) Validate() error {
	return c.validate()
}

func (c *Config) validate() error {
	for _, item := range c.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

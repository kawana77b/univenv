package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kawana77b/univenv/internal/sysutil"
)

type ConfigOpener struct {
	target string
	file   *sysutil.File
}

func (m *ConfigOpener) SetTarget(target string) {
	m.target = target
}

func (m *ConfigOpener) SetFile(path string) {
	if len(path) == 0 {
		return
	}
	m.setFile(sysutil.NewFile(path))
}

func (m *ConfigOpener) setFile(file *sysutil.File) {
	m.file = file
}

func (m *ConfigOpener) Dir() string {
	d := os.Getenv(ENV_UNIVENV_CONFIG_DIR)
	if d != "" {
		return d
	}
	return DEFAULT_DIR
}

func (m *ConfigOpener) FilePath() string {
	if m.file != nil {
		return m.file.Path()
	}

	createFilename := func(target string) string {
		basename := DEFAULT_FILE_BASENAME
		if target != "" {
			basename += fmt.Sprintf(".%s", target)
		}
		return fmt.Sprintf("%s%s", basename, DEFAULT_FILE_EXT)
	}

	dir := m.Dir()
	filename := createFilename(m.target)
	return filepath.Join(dir, filename)
}

func (m *ConfigOpener) Open() (*Config, error) {
	if m.file != nil {
		return m.open()
	}
	m.SetFile(m.FilePath())
	return m.open()
}

func (m *ConfigOpener) open() (*Config, error) {
	bytes, err := m.file.Read()
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = config.Read(bytes)
	if err != nil {
		return nil, err
	}
	return config, nil
}

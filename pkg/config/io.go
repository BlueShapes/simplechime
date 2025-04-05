package config

import (
	"bytes"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/blueshapes/simplechime/pkg/fileutil"
)

func doLoadConfig(configPath string) (*Config, error) {
	// Open file
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read all
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Decode config
	var config Config
	err = toml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	// Return result
	return &config, nil
}

func LoadConfig(configPath string) (*Config, error) {
	if !fileutil.FileExists(configPath) {
		err := SaveConfig(configPath, GetDefaultConfig())
		if err != nil {
			return nil, err
		}
	}
	return LoadConfig(configPath)
}

func SaveConfig(configPath string, config Config) error {
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	// copy all
	_, err = io.Copy(f, bytes.NewReader(b))
	return err
}

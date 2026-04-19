package config

import (
	"os"
	"path/filepath"
)

const (
	DefaultConfigDir   = ".harry-potter"
	DefaultConfigFile  = "config.yaml"
	DefaultDataDir     = "/var/db/harry-potter"
	DefaultLogDir      = "/var/log/harry-potter"
	DefaultTemplateDir = "/var/db/harry-potter/templates"
	DefaultZFSPool     = "zroot"
	DefaultZFSDataset  = "zroot/jails"
)

type Config struct {
	DataDir   string         `yaml:"data_dir"`
	LogDir    string         `yaml:"log_dir"`
	ZFS       ZFSConfig      `yaml:"zfs"`
	Network   NetworkConfig  `yaml:"network"`
	Templates TemplateConfig `yaml:"templates"`
}

type ZFSConfig struct {
	Enabled bool   `yaml:"enabled"`
	Pool    string `yaml:"pool"`
	Dataset string `yaml:"dataset"`
}

type NetworkConfig struct {
	DefaultInterface string `yaml:"default_interface"`
	DefaultSubnet    string `yaml:"default_subnet"`
}

type TemplateConfig struct {
	Dir    string `yaml:"dir"`
	Mirror string `yaml:"mirror"`
}

func Default() *Config {
	return &Config{
		DataDir: DefaultDataDir,
		LogDir:  DefaultLogDir,
		ZFS: ZFSConfig{
			Enabled: true,
			Pool:    DefaultZFSPool,
			Dataset: DefaultZFSDataset,
		},
		Network: NetworkConfig{
			DefaultInterface: "em0",
			DefaultSubnet:    "10.0.0.0/24",
		},
		Templates: TemplateConfig{
			Dir:    DefaultTemplateDir,
			Mirror: "https://download.freebsd.org/releases",
		},
	}
}

func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, DefaultConfigDir), nil
}

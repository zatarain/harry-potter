package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfigDir(t *testing.T) {
	if DefaultConfigDir == "" {
		t.Error("DefaultConfigDir must not be empty")
	}
}

func TestDefaultConstants(t *testing.T) {
	cases := map[string]string{
		"DefaultConfigDir":   DefaultConfigDir,
		"DefaultConfigFile":  DefaultConfigFile,
		"DefaultDataDir":     DefaultDataDir,
		"DefaultLogDir":      DefaultLogDir,
		"DefaultTemplateDir": DefaultTemplateDir,
		"DefaultZFSPool":     DefaultZFSPool,
		"DefaultZFSDataset":  DefaultZFSDataset,
	}
	for name, value := range cases {
		if value == "" {
			t.Errorf("constant %s must not be empty", name)
		}
	}
}

func TestDefaultReturnsNonNil(t *testing.T) {
	cfg := Default()
	if cfg == nil {
		t.Fatal("Default() returned nil")
	}
}

func TestDefaultDataDir(t *testing.T) {
	cfg := Default()
	if cfg.DataDir != DefaultDataDir {
		t.Errorf("Default().DataDir = %q, want %q", cfg.DataDir, DefaultDataDir)
	}
}

func TestDefaultLogDir(t *testing.T) {
	cfg := Default()
	if cfg.LogDir != DefaultLogDir {
		t.Errorf("Default().LogDir = %q, want %q", cfg.LogDir, DefaultLogDir)
	}
}

func TestDefaultZFSEnabled(t *testing.T) {
	cfg := Default()
	if !cfg.ZFS.Enabled {
		t.Error("Default().ZFS.Enabled should be true")
	}
}

func TestDefaultZFSPool(t *testing.T) {
	cfg := Default()
	if cfg.ZFS.Pool != DefaultZFSPool {
		t.Errorf("Default().ZFS.Pool = %q, want %q", cfg.ZFS.Pool, DefaultZFSPool)
	}
}

func TestDefaultZFSDataset(t *testing.T) {
	cfg := Default()
	if cfg.ZFS.Dataset != DefaultZFSDataset {
		t.Errorf("Default().ZFS.Dataset = %q, want %q", cfg.ZFS.Dataset, DefaultZFSDataset)
	}
}

func TestDefaultNetworkSubnet(t *testing.T) {
	cfg := Default()
	if cfg.Network.DefaultSubnet == "" {
		t.Error("Default().Network.DefaultSubnet must not be empty")
	}
}

func TestDefaultNetworkInterface(t *testing.T) {
	cfg := Default()
	if cfg.Network.DefaultInterface == "" {
		t.Error("Default().Network.DefaultInterface must not be empty")
	}
}

func TestDefaultTemplateMirror(t *testing.T) {
	cfg := Default()
	if cfg.Templates.Mirror == "" {
		t.Error("Default().Templates.Mirror must not be empty")
	}
}

func TestDir(t *testing.T) {
	dir, err := Dir()
	if err != nil {
		t.Fatalf("Dir() returned error: %v", err)
	}
	if dir == "" {
		t.Fatal("Dir() returned empty string")
	}
	home, _ := os.UserHomeDir()
	want := filepath.Join(home, DefaultConfigDir)
	if dir != want {
		t.Errorf("Dir() = %q, want %q", dir, want)
	}
}

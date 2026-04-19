package command

import (
	"strings"
	"testing"
)

func TestVersionBuildVariableDefaults(t *testing.T) {
	for name, value := range map[string]string{
		"Version":   Version,
		"Commit":    Commit,
		"BuildDate": BuildDate,
	} {
		if value == "" {
			t.Errorf("build variable %s should not be empty", name)
		}
	}
}

func TestVersionCommandUse(t *testing.T) {
	if !strings.HasPrefix(versionCommand.Use, "version") {
		t.Errorf("versionCommand.Use = %q, want prefix 'version'", versionCommand.Use)
	}
}

func TestVersionCommandShortDescription(t *testing.T) {
	if versionCommand.Short == "" {
		t.Error("versionCommand.Short should not be empty")
	}
}

func TestVersionCommandHasShortFlag(t *testing.T) {
	if !hasFlag(versionCommand, "short") {
		t.Error("--short flag not registered on versionCommand")
	}
}

package command

import (
	"strings"
	"testing"
)

func TestSystemCommandUse(t *testing.T) {
	if systemCommand.Use != "system" {
		t.Errorf("systemCommand.Use = %q, want 'system'", systemCommand.Use)
	}
}

func TestSystemSubcommands(t *testing.T) {
	for _, name := range []string{"df", "events", "info", "prune"} {
		if !hasSubcommand(systemCommand, name) {
			t.Errorf("subcommand %q not found on systemCommand", name)
		}
	}
}

func TestSystemDiskUsageCommandUse(t *testing.T) {
	if !strings.HasPrefix(systemDiskUsageCommand.Use, "df") {
		t.Errorf("systemDiskUsageCommand.Use = %q, want prefix 'df'", systemDiskUsageCommand.Use)
	}
}

func TestSystemDiskUsageFlags(t *testing.T) {
	for _, flag := range []string{"format", "verbose"} {
		if !hasFlag(systemDiskUsageCommand, flag) {
			t.Errorf("flag --%s not registered on systemDiskUsageCommand", flag)
		}
	}
}

func TestSystemEventsHasAlias(t *testing.T) {
	if !hasAlias(systemEventsCommand, "ev") {
		t.Error("alias 'ev' not found in systemEventsCommand.Aliases")
	}
}

func TestSystemEventsFlags(t *testing.T) {
	for _, flag := range []string{"filter", "format", "since", "until"} {
		if !hasFlag(systemEventsCommand, flag) {
			t.Errorf("flag --%s not registered on systemEventsCommand", flag)
		}
	}
}

func TestSystemPruneFlags(t *testing.T) {
	for _, flag := range []string{"all", "filter", "force", "volumes"} {
		if !hasFlag(systemPruneCommand, flag) {
			t.Errorf("flag --%s not registered on systemPruneCommand", flag)
		}
	}
}

func TestSystemInfoFlags(t *testing.T) {
	if !hasFlag(systemInfoCommand, "format") {
		t.Error("flag --format not registered on systemInfoCommand")
	}
}

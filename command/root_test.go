package command

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// ── shared test helpers ───────────────────────────────────────────────────────

func hasSubcommand(parent *cobra.Command, name string) bool {
	for _, sub := range parent.Commands() {
		if sub.Name() == name {
			return true
		}
	}
	return false
}

func hasFlag(target *cobra.Command, name string) bool {
	return target.Flags().Lookup(name) != nil
}

func hasPersistentFlag(target *cobra.Command, name string) bool {
	return target.PersistentFlags().Lookup(name) != nil
}

func hasAlias(target *cobra.Command, alias string) bool {
	for _, a := range target.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestNotImplemented(t *testing.T) {
	sentinel := &cobra.Command{Use: "sentinel"}
	err := notImplemented(sentinel, nil)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if !strings.Contains(err.Error(), "not yet implemented") {
		t.Errorf("error %q missing phrase 'not yet implemented'", err.Error())
	}
	if !strings.Contains(err.Error(), "sentinel") {
		t.Errorf("error %q missing the command name 'sentinel'", err.Error())
	}
}

func TestBinaryName(t *testing.T) {
	if binaryName() == "" {
		t.Error("binaryName() returned an empty string")
	}
}

func TestRootCommandPersistentFlags(t *testing.T) {
	for _, flag := range []string{"config", "verbose"} {
		if !hasPersistentFlag(rootCommand, flag) {
			t.Errorf("persistent flag --%s not registered on rootCommand", flag)
		}
	}
}

func TestRootCommandManagementSubcommands(t *testing.T) {
	for _, name := range []string{"jail", "image", "network", "volume", "system", "version"} {
		if !hasSubcommand(rootCommand, name) {
			t.Errorf("management subcommand %q not found on rootCommand", name)
		}
	}
}

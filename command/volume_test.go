package command

import "testing"

func TestVolumeCommandUse(t *testing.T) {
	if volumeCommand.Use != "volume" {
		t.Errorf("volumeCommand.Use = %q, want 'volume'", volumeCommand.Use)
	}
}

func TestVolumeCommandHasVolAlias(t *testing.T) {
	if !hasAlias(volumeCommand, "vol") {
		t.Error("alias 'vol' not found in volumeCommand.Aliases")
	}
}

func TestVolumeSubcommands(t *testing.T) {
	for _, name := range []string{"create", "inspect", "ls", "prune", "rm"} {
		if !hasSubcommand(volumeCommand, name) {
			t.Errorf("subcommand %q not found on volumeCommand", name)
		}
	}
}

func TestVolumeListAliases(t *testing.T) {
	if !hasAlias(volumeListCommand, "list") {
		t.Error("alias 'list' not found in volumeListCommand.Aliases")
	}
}

func TestVolumeRemoveAliases(t *testing.T) {
	if !hasAlias(volumeRemoveCommand, "remove") {
		t.Error("alias 'remove' not found in volumeRemoveCommand.Aliases")
	}
}

func TestVolumeCreateFlags(t *testing.T) {
	for _, flag := range []string{"driver", "label", "opt"} {
		if !hasFlag(volumeCreateCommand, flag) {
			t.Errorf("flag --%s not registered on volumeCreateCommand", flag)
		}
	}
}

func TestVolumeListFlags(t *testing.T) {
	for _, flag := range []string{"cluster", "filter", "format", "quiet"} {
		if !hasFlag(volumeListCommand, flag) {
			t.Errorf("flag --%s not registered on volumeListCommand", flag)
		}
	}
}

func TestVolumePruneFlags(t *testing.T) {
	for _, flag := range []string{"all", "filter", "force"} {
		if !hasFlag(volumePruneCommand, flag) {
			t.Errorf("flag --%s not registered on volumePruneCommand", flag)
		}
	}
}

func TestVolumeRemoveFlags(t *testing.T) {
	if !hasFlag(volumeRemoveCommand, "force") {
		t.Error("flag --force not registered on volumeRemoveCommand")
	}
}

func TestVolumeInspectFlags(t *testing.T) {
	if !hasFlag(volumeInspectCommand, "format") {
		t.Error("flag --format not registered on volumeInspectCommand")
	}
}

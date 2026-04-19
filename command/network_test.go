package command

import "testing"

func TestNetworkCommandUse(t *testing.T) {
	if networkCommand.Use != "network" {
		t.Errorf("networkCommand.Use = %q, want 'network'", networkCommand.Use)
	}
}

func TestNetworkCommandHasNetAlias(t *testing.T) {
	if !hasAlias(networkCommand, "net") {
		t.Error("alias 'net' not found in networkCommand.Aliases")
	}
}

func TestNetworkSubcommands(t *testing.T) {
	for _, name := range []string{"connect", "create", "disconnect", "inspect", "ls", "prune", "rm"} {
		if !hasSubcommand(networkCommand, name) {
			t.Errorf("subcommand %q not found on networkCommand", name)
		}
	}
}

func TestNetworkRemoveAliases(t *testing.T) {
	if !hasAlias(networkRemoveCommand, "remove") {
		t.Error("alias 'remove' not found in networkRemoveCommand.Aliases")
	}
}

func TestNetworkListAliases(t *testing.T) {
	if !hasAlias(networkListCommand, "list") {
		t.Error("alias 'list' not found in networkListCommand.Aliases")
	}
}

func TestNetworkCreateFlags(t *testing.T) {
	for _, flag := range []string{"driver", "gateway", "internal", "ipv6", "label", "opt", "scope", "subnet"} {
		if !hasFlag(networkCreateCommand, flag) {
			t.Errorf("flag --%s not registered on networkCreateCommand", flag)
		}
	}
}

func TestNetworkConnectFlags(t *testing.T) {
	for _, flag := range []string{"alias", "ip", "ip6", "link"} {
		if !hasFlag(networkConnectCommand, flag) {
			t.Errorf("flag --%s not registered on networkConnectCommand", flag)
		}
	}
}

func TestNetworkListFlags(t *testing.T) {
	for _, flag := range []string{"filter", "format", "no-trunc", "quiet"} {
		if !hasFlag(networkListCommand, flag) {
			t.Errorf("flag --%s not registered on networkListCommand", flag)
		}
	}
}

func TestNetworkDisconnectFlags(t *testing.T) {
	if !hasFlag(networkDisconnectCommand, "force") {
		t.Error("flag --force not registered on networkDisconnectCommand")
	}
}

func TestNetworkInspectFlags(t *testing.T) {
	for _, flag := range []string{"format", "verbose"} {
		if !hasFlag(networkInspectCommand, flag) {
			t.Errorf("flag --%s not registered on networkInspectCommand", flag)
		}
	}
}

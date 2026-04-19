package command

import "testing"

func TestJailCommandUse(t *testing.T) {
	if jailCommand.Use != "jail" {
		t.Errorf("jailCommand.Use = %q, want 'jail'", jailCommand.Use)
	}
}

func TestJailCommandAliases(t *testing.T) {
	for _, alias := range []string{"container", "c", "pot", "p", "j"} {
		if !hasAlias(jailCommand, alias) {
			t.Errorf("alias %q not found in jailCommand.Aliases", alias)
		}
	}
}

func TestJailSubcommands(t *testing.T) {
	for _, name := range []string{
		"attach", "commit", "cp", "create", "diff", "exec", "export",
		"inspect", "kill", "logs", "ls", "pause", "port", "prune",
		"rename", "restart", "rm", "run", "start", "stats", "stop",
		"top", "unpause", "update", "wait",
	} {
		if !hasSubcommand(jailCommand, name) {
			t.Errorf("subcommand %q not found on jailCommand", name)
		}
	}
}

func TestJailListCommandAliases(t *testing.T) {
	for _, alias := range []string{"ps", "list"} {
		if !hasAlias(jailListCommand, alias) {
			t.Errorf("alias %q not found in jailListCommand.Aliases", alias)
		}
	}
}

func TestJailRunFlags(t *testing.T) {
	for _, flag := range []string{
		"detach", "env", "name", "network", "publish", "tty", "interactive",
		"volume", "rm", "memory", "cpus", "hostname", "restart", "user",
		"workdir", "privileged", "ip", "label", "entrypoint", "pull",
	} {
		if !hasFlag(jailRunCommand, flag) {
			t.Errorf("flag --%s not registered on jailRunCommand", flag)
		}
	}
}

func TestJailCreateSharesRunFlags(t *testing.T) {
	for _, flag := range []string{"name", "network", "volume", "memory", "cpus"} {
		if !hasFlag(jailCreateCommand, flag) {
			t.Errorf("flag --%s not registered on jailCreateCommand", flag)
		}
	}
}

func TestJailListFlags(t *testing.T) {
	for _, flag := range []string{"all", "filter", "format", "quiet", "size", "no-trunc", "last", "latest"} {
		if !hasFlag(jailListCommand, flag) {
			t.Errorf("flag --%s not registered on jailListCommand", flag)
		}
	}
}

func TestJailLogsFlags(t *testing.T) {
	for _, flag := range []string{"follow", "tail", "since", "until", "timestamps", "details"} {
		if !hasFlag(jailLogsCommand, flag) {
			t.Errorf("flag --%s not registered on jailLogsCommand", flag)
		}
	}
}

func TestJailExecFlags(t *testing.T) {
	for _, flag := range []string{"detach", "env", "interactive", "tty", "user", "workdir", "privileged"} {
		if !hasFlag(jailExecCommand, flag) {
			t.Errorf("flag --%s not registered on jailExecCommand", flag)
		}
	}
}

func TestJailStopFlags(t *testing.T) {
	for _, flag := range []string{"signal", "time"} {
		if !hasFlag(jailStopCommand, flag) {
			t.Errorf("flag --%s not registered on jailStopCommand", flag)
		}
	}
}

func TestJailRemoveFlags(t *testing.T) {
	for _, flag := range []string{"force", "volumes"} {
		if !hasFlag(jailRemoveCommand, flag) {
			t.Errorf("flag --%s not registered on jailRemoveCommand", flag)
		}
	}
}

func TestJailAttachFlags(t *testing.T) {
	for _, flag := range []string{"detach-keys", "no-stdin", "sig-proxy"} {
		if !hasFlag(jailAttachCommand, flag) {
			t.Errorf("flag --%s not registered on jailAttachCommand", flag)
		}
	}
}

func TestJailCommitFlags(t *testing.T) {
	for _, flag := range []string{"author", "change", "message", "pause"} {
		if !hasFlag(jailCommitCommand, flag) {
			t.Errorf("flag --%s not registered on jailCommitCommand", flag)
		}
	}
}

func TestJailStatsFlags(t *testing.T) {
	for _, flag := range []string{"all", "format", "no-stream", "no-trunc"} {
		if !hasFlag(jailStatsCommand, flag) {
			t.Errorf("flag --%s not registered on jailStatsCommand", flag)
		}
	}
}

func TestJailUpdateFlags(t *testing.T) {
	for _, flag := range []string{"cpus", "memory", "memory-swap", "pids-limit", "restart"} {
		if !hasFlag(jailUpdateCommand, flag) {
			t.Errorf("flag --%s not registered on jailUpdateCommand", flag)
		}
	}
}

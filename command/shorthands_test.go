package command

import (
	"strings"
	"testing"
)

func TestJailShorthandsRegisteredOnRoot(t *testing.T) {
	for _, name := range []string{
		"attach", "commit", "cp", "create", "diff", "events", "exec", "export",
		"kill", "logs", "pause", "port", "ps", "rename", "restart", "rm", "run",
		"start", "stats", "stop", "top", "unpause", "update", "wait",
	} {
		if !hasSubcommand(rootCommand, name) {
			t.Errorf("jail shorthand %q not registered on rootCommand", name)
		}
	}
}

func TestImageShorthandsRegisteredOnRoot(t *testing.T) {
	for _, name := range []string{
		"build", "history", "images", "import", "load", "pull", "push", "rmi", "save", "tag",
	} {
		if !hasSubcommand(rootCommand, name) {
			t.Errorf("image shorthand %q not registered on rootCommand", name)
		}
	}
}

func TestSystemShorthandsRegisteredOnRoot(t *testing.T) {
	for _, name := range []string{"info", "inspect", "login", "logout", "search"} {
		if !hasSubcommand(rootCommand, name) {
			t.Errorf("system shorthand %q not registered on rootCommand", name)
		}
	}
}

func TestCopyCommandUse(t *testing.T) {
	if !strings.HasPrefix(copyCommand.Use, "cp") {
		t.Errorf("copyCommand.Use = %q, want prefix 'cp'", copyCommand.Use)
	}
}

func TestPsCommandUse(t *testing.T) {
	if !strings.HasPrefix(psCommand.Use, "ps") {
		t.Errorf("psCommand.Use = %q, want prefix 'ps'", psCommand.Use)
	}
}

func TestRemoveCommandUse(t *testing.T) {
	if !strings.HasPrefix(removeCommand.Use, "rm") {
		t.Errorf("removeCommand.Use = %q, want prefix 'rm'", removeCommand.Use)
	}
}

func TestRemoveImageCommandUse(t *testing.T) {
	if !strings.HasPrefix(removeImageCommand.Use, "rmi") {
		t.Errorf("removeImageCommand.Use = %q, want prefix 'rmi'", removeImageCommand.Use)
	}
}

func TestRunCommandFlags(t *testing.T) {
	for _, flag := range []string{"detach", "env", "name", "network", "publish", "tty", "interactive",
		"volume", "rm", "memory", "cpus", "hostname", "restart", "user", "workdir", "privileged", "ip",
		"label", "entrypoint", "pull"} {
		if !hasFlag(runCommand, flag) {
			t.Errorf("flag --%s not registered on runCommand", flag)
		}
	}
}

func TestPsCommandFlags(t *testing.T) {
	for _, flag := range []string{"all", "filter", "format", "quiet", "size", "no-trunc", "last", "latest"} {
		if !hasFlag(psCommand, flag) {
			t.Errorf("flag --%s not registered on psCommand", flag)
		}
	}
}

func TestLogsCommandFlags(t *testing.T) {
	for _, flag := range []string{"follow", "tail", "since", "until", "timestamps", "details"} {
		if !hasFlag(logsCommand, flag) {
			t.Errorf("flag --%s not registered on logsCommand", flag)
		}
	}
}

func TestExecCommandFlags(t *testing.T) {
	for _, flag := range []string{"detach", "env", "interactive", "tty", "user", "workdir", "privileged"} {
		if !hasFlag(execCommand, flag) {
			t.Errorf("flag --%s not registered on execCommand", flag)
		}
	}
}

func TestAttachCommandFlags(t *testing.T) {
	for _, flag := range []string{"detach-keys", "no-stdin", "sig-proxy"} {
		if !hasFlag(attachCommand, flag) {
			t.Errorf("flag --%s not registered on attachCommand", flag)
		}
	}
}

func TestStopCommandFlags(t *testing.T) {
	for _, flag := range []string{"signal", "time"} {
		if !hasFlag(stopCommand, flag) {
			t.Errorf("flag --%s not registered on stopCommand", flag)
		}
	}
}

func TestRemoveCommandFlags(t *testing.T) {
	for _, flag := range []string{"force", "volumes"} {
		if !hasFlag(removeCommand, flag) {
			t.Errorf("flag --%s not registered on removeCommand", flag)
		}
	}
}

func TestBuildCommandFlags(t *testing.T) {
	for _, flag := range []string{"file", "tag", "no-cache", "build-arg", "platform", "pull", "quiet", "target"} {
		if !hasFlag(buildCommand, flag) {
			t.Errorf("flag --%s not registered on buildCommand", flag)
		}
	}
}

func TestImagesCommandFlags(t *testing.T) {
	for _, flag := range []string{"all", "digests", "filter", "format", "no-trunc", "quiet"} {
		if !hasFlag(imagesCommand, flag) {
			t.Errorf("flag --%s not registered on imagesCommand", flag)
		}
	}
}

func TestPullCommandFlags(t *testing.T) {
	for _, flag := range []string{"all-tags", "disable-content-trust", "platform", "quiet"} {
		if !hasFlag(pullCommand, flag) {
			t.Errorf("flag --%s not registered on pullCommand", flag)
		}
	}
}

func TestPushCommandFlags(t *testing.T) {
	for _, flag := range []string{"all-tags", "disable-content-trust", "quiet"} {
		if !hasFlag(pushCommand, flag) {
			t.Errorf("flag --%s not registered on pushCommand", flag)
		}
	}
}

package command

import "github.com/spf13/cobra"

var jailCommand = &cobra.Command{
	Use:     "jail",
	Short:   "Manage jails",
	Aliases: []string{"j", "container", "c", "pot", "p"},
}

// ── subcommands ──────────────────────────────────────────────────────────────

var jailAttachCommand = &cobra.Command{
	Use:   "attach [OPTIONS] JAIL",
	Short: "Attach local standard input, output, and error streams to a running jail",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var jailCommitCommand = &cobra.Command{
	Use:   "commit [OPTIONS] JAIL [REPOSITORY[:TAG]]",
	Short: "Create a new image from a jail's changes",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var jailCopyCommand = &cobra.Command{
	Use:   "cp [OPTIONS] JAIL:SRC_PATH DEST_PATH|-\tcp [OPTIONS] SRC_PATH|- JAIL:DEST_PATH",
	Short: "Copy files/folders between a jail and the local filesystem",
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var jailCreateCommand = &cobra.Command{
	Use:   "create [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Create a new jail without starting it",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailDiffCommand = &cobra.Command{
	Use:   "diff JAIL",
	Short: "Inspect changes to files or directories on a jail's filesystem",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var jailExecCommand = &cobra.Command{
	Use:   "exec [OPTIONS] JAIL COMMAND [ARG...]",
	Short: "Execute a command in a running jail",
	Args:  cobra.MinimumNArgs(2),
	RunE:  notImplemented,
}

var jailExportCommand = &cobra.Command{
	Use:   "export [OPTIONS] JAIL",
	Short: "Export a jail's filesystem as a tar archive",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var jailInspectCommand = &cobra.Command{
	Use:   "inspect [OPTIONS] JAIL [JAIL...]",
	Short: "Display detailed information on one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailKillCommand = &cobra.Command{
	Use:   "kill [OPTIONS] JAIL [JAIL...]",
	Short: "Kill one or more running jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailLogsCommand = &cobra.Command{
	Use:   "logs [OPTIONS] JAIL",
	Short: "Fetch the logs of a jail",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var jailListCommand = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Aliases: []string{"ps", "list"},
	Short:   "List jails",
	RunE:    notImplemented,
}

var jailPauseCommand = &cobra.Command{
	Use:   "pause JAIL [JAIL...]",
	Short: "Pause all processes within one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailPortCommand = &cobra.Command{
	Use:   "port JAIL [PRIVATE_PORT[/PROTO]]",
	Short: "List port mappings or a specific mapping for the jail",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var jailPruneCommand = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove all stopped jails",
	RunE:  notImplemented,
}

var jailRenameCommand = &cobra.Command{
	Use:   "rename JAIL NEW_NAME",
	Short: "Rename a jail",
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var jailRestartCommand = &cobra.Command{
	Use:   "restart [OPTIONS] JAIL [JAIL...]",
	Short: "Restart one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailRemoveCommand = &cobra.Command{
	Use:   "rm [OPTIONS] JAIL [JAIL...]",
	Short: "Remove one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailRunCommand = &cobra.Command{
	Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Create and run a new jail from an image",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailStartCommand = &cobra.Command{
	Use:   "start [OPTIONS] JAIL [JAIL...]",
	Short: "Start one or more stopped jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailStatsCommand = &cobra.Command{
	Use:   "stats [OPTIONS] [JAIL...]",
	Short: "Display a live stream of jail resource usage statistics",
	RunE:  notImplemented,
}

var jailStopCommand = &cobra.Command{
	Use:   "stop [OPTIONS] JAIL [JAIL...]",
	Short: "Stop one or more running jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailTopCommand = &cobra.Command{
	Use:   "top JAIL [ps OPTIONS]",
	Short: "Display the running processes of a jail",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailUnpauseCommand = &cobra.Command{
	Use:   "unpause JAIL [JAIL...]",
	Short: "Unpause all processes within one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailUpdateCommand = &cobra.Command{
	Use:   "update [OPTIONS] JAIL [JAIL...]",
	Short: "Update configuration of one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var jailWaitCommand = &cobra.Command{
	Use:   "wait JAIL [JAIL...]",
	Short: "Block until one or more jails stop, then print their exit codes",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

// ── flag helpers (shared with top-level shorthands) ──────────────────────────

func addJailRunFlags(target *cobra.Command) {
	target.Flags().StringSliceP("attach", "a", nil, "Attach to STDIN, STDOUT or STDERR")
	target.Flags().StringSlice("cap-add", nil, "Add jail capabilities/parameters")
	target.Flags().StringSlice("cap-drop", nil, "Drop jail capabilities/parameters")
	target.Flags().String("cidfile", "", "Write the jail ID to this file")
	target.Flags().String("cpus", "", "Number of CPUs (RCTL limit)")
	target.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	target.Flags().BoolP("detach", "d", false, "Run jail in the background")
	target.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	target.Flags().StringSliceP("env", "e", nil, "Set environment variables")
	target.Flags().String("env-file", "", "Read environment variables from a file")
	target.Flags().String("entrypoint", "", "Overwrite the default entrypoint of the image")
	target.Flags().String("hostname", "", "Jail hostname")
	target.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	target.Flags().String("ip", "", "IPv4 address")
	target.Flags().String("ip6", "", "IPv6 address")
	target.Flags().StringSliceP("label", "l", nil, "Set metadata on a jail")
	target.Flags().String("label-file", "", "Read labels from a file")
	target.Flags().StringP("memory", "m", "", "Memory limit (RCTL, e.g. 512m, 2g)")
	target.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap")
	target.Flags().String("name", "", "Assign a name to the jail")
	target.Flags().String("network", "", "Connect a jail to a network")
	target.Flags().StringSlice("network-alias", nil, "Add network-scoped alias for the jail")
	target.Flags().Bool("privileged", false, "Give extended privileges to this jail")
	target.Flags().StringSliceP("publish", "p", nil, "Publish jail ports to the host (via pf(4))")
	target.Flags().BoolP("publish-all", "P", false, "Publish all exposed ports to random host ports")
	target.Flags().String("pull", "missing", "Pull image before running (always, missing, never)")
	target.Flags().BoolP("quiet", "q", false, "Suppress the pull output")
	target.Flags().Bool("read-only", false, "Mount the jail's root filesystem as read only")
	target.Flags().String("restart", "no", "Restart policy (no, always, on-failure[:n], unless-stopped)")
	target.Flags().Bool("rm", false, "Automatically remove the jail and its image when it exits")
	target.Flags().String("stop-signal", "SIGTERM", "Signal to stop the jail")
	target.Flags().Int("stop-timeout", 10, "Timeout (seconds) to stop a jail before killing it")
	target.Flags().StringToString("sysctl", nil, "Sysctl options")
	target.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	target.Flags().StringP("user", "u", "", "Username or UID inside the jail")
	target.Flags().StringSliceP("volume", "v", nil, "Bind mount a volume (nullfs)")
	target.Flags().StringSlice("volumes-from", nil, "Mount volumes from the specified jail(s)")
	target.Flags().StringP("workdir", "w", "", "Working directory inside the jail")
}

func addJailListFlags(target *cobra.Command) {
	target.Flags().BoolP("all", "a", false, "Show all jails (default shows just running)")
	target.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	target.Flags().String("format", "", "Format output using a custom template:\n'table': print output in table format with column headers (default)\n'table TEMPLATE': print output in table format using the given Go template\n'json': print in JSON format\n'TEMPLATE': print output using the given Go template")
	target.Flags().IntP("last", "n", -1, "Show n last created jails (includes all states)")
	target.Flags().BoolP("latest", "l", false, "Show the latest created jail (includes all states)")
	target.Flags().Bool("no-trunc", false, "Don't truncate output")
	target.Flags().BoolP("quiet", "q", false, "Only display jail IDs")
	target.Flags().BoolP("size", "s", false, "Display total file sizes")
}

func addJailLogsFlags(target *cobra.Command) {
	target.Flags().Bool("details", false, "Show extra details provided to logs")
	target.Flags().BoolP("follow", "f", false, "Follow log output")
	target.Flags().String("since", "", "Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m)")
	target.Flags().StringP("tail", "n", "all", "Number of lines to show from the end of the logs")
	target.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	target.Flags().String("until", "", "Show logs before a timestamp or relative")
}

func addJailExecFlags(target *cobra.Command) {
	target.Flags().BoolP("detach", "d", false, "Detached mode: run command in the background")
	target.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	target.Flags().StringSliceP("env", "e", nil, "Set environment variables")
	target.Flags().String("env-file", "", "Read from a file of environment variables")
	target.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	target.Flags().Bool("privileged", false, "Give extended privileges to the command")
	target.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	target.Flags().StringP("user", "u", "", "Username or UID")
	target.Flags().StringP("workdir", "w", "", "Working directory inside the jail")
}

func addJailStopFlags(target *cobra.Command) {
	target.Flags().StringP("signal", "s", "", "Signal to send to the jail")
	target.Flags().IntP("time", "t", 10, "Seconds to wait for stop before killing")
}

func addJailRemoveFlags(target *cobra.Command) {
	target.Flags().BoolP("force", "f", false, "Force the removal of a running jail")
	target.Flags().BoolP("volumes", "v", false, "Remove anonymous volumes associated with the jail")
}

func init() {
	addJailRunFlags(jailRunCommand)
	addJailRunFlags(jailCreateCommand)

	addJailListFlags(jailListCommand)

	addJailLogsFlags(jailLogsCommand)

	addJailExecFlags(jailExecCommand)

	addJailStopFlags(jailStopCommand)

	addJailRemoveFlags(jailRemoveCommand)

	jailAttachCommand.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	jailAttachCommand.Flags().Bool("no-stdin", false, "Do not attach STDIN")
	jailAttachCommand.Flags().Bool("sig-proxy", true, "Proxy all received signals to the process")

	jailCommitCommand.Flags().StringP("author", "a", "", "Author (e.g., 'John Smith <john@example.com>')")
	jailCommitCommand.Flags().StringSliceP("change", "c", nil, "Apply a Jailfile instruction to the created image")
	jailCommitCommand.Flags().StringP("message", "m", "", "Commit message")
	jailCommitCommand.Flags().BoolP("pause", "p", true, "Pause jail during commit")

	jailCopyCommand.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	jailCopyCommand.Flags().BoolP("follow-link", "L", false, "Always follow symbolic link in SRC_PATH")
	jailCopyCommand.Flags().BoolP("quiet", "q", false, "Suppress progress output during copy")

	jailExportCommand.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")

	jailInspectCommand.Flags().StringP("format", "f", "", "Format output using a custom template")
	jailInspectCommand.Flags().BoolP("size", "s", false, "Display total file sizes")

	jailKillCommand.Flags().StringP("signal", "s", "KILL", "Signal to send to the jail")

	jailPruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	jailPruneCommand.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")

	jailRestartCommand.Flags().StringP("signal", "s", "", "Signal to send to the jail")
	jailRestartCommand.Flags().IntP("time", "t", 10, "Seconds to wait for stop before killing")

	jailStartCommand.Flags().BoolP("attach", "a", false, "Attach STDOUT/STDERR and forward signals")
	jailStartCommand.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	jailStartCommand.Flags().BoolP("interactive", "i", false, "Attach jail's STDIN")

	jailStatsCommand.Flags().BoolP("all", "a", false, "Show all jails (default shows just running)")
	jailStatsCommand.Flags().String("format", "", "Format output using a custom template")
	jailStatsCommand.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	jailStatsCommand.Flags().Bool("no-trunc", false, "Do not truncate output")

	jailUpdateCommand.Flags().String("cpus", "", "Number of CPUs (RCTL limit)")
	jailUpdateCommand.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution")
	jailUpdateCommand.Flags().StringP("memory", "m", "", "Memory limit")
	jailUpdateCommand.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap")
	jailUpdateCommand.Flags().String("pids-limit", "", "Tune jail process limits via RCTL")
	jailUpdateCommand.Flags().String("restart", "", "Restart policy to apply")

	jailCommand.AddCommand(
		jailAttachCommand,
		jailCommitCommand,
		jailCopyCommand,
		jailCreateCommand,
		jailDiffCommand,
		jailExecCommand,
		jailExportCommand,
		jailInspectCommand,
		jailKillCommand,
		jailLogsCommand,
		jailListCommand,
		jailPauseCommand,
		jailPortCommand,
		jailPruneCommand,
		jailRenameCommand,
		jailRestartCommand,
		jailRemoveCommand,
		jailRunCommand,
		jailStartCommand,
		jailStatsCommand,
		jailStopCommand,
		jailTopCommand,
		jailUnpauseCommand,
		jailUpdateCommand,
		jailWaitCommand,
	)
	rootCommand.AddCommand(jailCommand)
}

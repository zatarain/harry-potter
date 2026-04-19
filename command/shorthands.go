package command

// Top-level shorthand commands that mirror their management-command equivalents.
// Each delegates to the same RunE as the canonical subcommand and shares its
// flag set via the addJail*/addImage* helpers defined in jail.go and image.go.

import "github.com/spf13/cobra"

// ── jail shorthands ──────────────────────────────────────────────────────────

var attachCommand = &cobra.Command{
	Use:   "attach [OPTIONS] JAIL",
	Short: jailAttachCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var commitCommand = &cobra.Command{
	Use:   "commit [OPTIONS] JAIL [REPOSITORY[:TAG]]",
	Short: jailCommitCommand.Short,
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var copyCommand = &cobra.Command{
	Use:   "cp [OPTIONS] JAIL:SRC_PATH DEST_PATH|-\tcp [OPTIONS] SRC_PATH|- JAIL:DEST_PATH",
	Short: jailCopyCommand.Short,
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var createCommand = &cobra.Command{
	Use:   "create [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: jailCreateCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var diffCommand = &cobra.Command{
	Use:   "diff JAIL",
	Short: jailDiffCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var eventsCommand = &cobra.Command{
	Use:   "events [OPTIONS]",
	Short: systemEventsCommand.Short,
	RunE:  notImplemented,
}

var execCommand = &cobra.Command{
	Use:   "exec [OPTIONS] JAIL COMMAND [ARG...]",
	Short: jailExecCommand.Short,
	Args:  cobra.MinimumNArgs(2),
	RunE:  notImplemented,
}

var exportCommand = &cobra.Command{
	Use:   "export [OPTIONS] JAIL",
	Short: jailExportCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var killCommand = &cobra.Command{
	Use:   "kill [OPTIONS] JAIL [JAIL...]",
	Short: jailKillCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var logsCommand = &cobra.Command{
	Use:   "logs [OPTIONS] JAIL",
	Short: jailLogsCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var pauseCommand = &cobra.Command{
	Use:   "pause JAIL [JAIL...]",
	Short: jailPauseCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var portCommand = &cobra.Command{
	Use:   "port JAIL [PRIVATE_PORT[/PROTO]]",
	Short: jailPortCommand.Short,
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var psCommand = &cobra.Command{
	Use:   "ps [OPTIONS]",
	Short: jailListCommand.Short,
	RunE:  notImplemented,
}

var renameCommand = &cobra.Command{
	Use:   "rename JAIL NEW_NAME",
	Short: jailRenameCommand.Short,
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var restartCommand = &cobra.Command{
	Use:   "restart [OPTIONS] JAIL [JAIL...]",
	Short: jailRestartCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var removeCommand = &cobra.Command{
	Use:   "rm [OPTIONS] JAIL [JAIL...]",
	Short: jailRemoveCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var runCommand = &cobra.Command{
	Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: jailRunCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var startCommand = &cobra.Command{
	Use:   "start [OPTIONS] JAIL [JAIL...]",
	Short: jailStartCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var statsCommand = &cobra.Command{
	Use:   "stats [OPTIONS] [JAIL...]",
	Short: jailStatsCommand.Short,
	RunE:  notImplemented,
}

var stopCommand = &cobra.Command{
	Use:   "stop [OPTIONS] JAIL [JAIL...]",
	Short: jailStopCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var topCommand = &cobra.Command{
	Use:   "top JAIL [ps OPTIONS]",
	Short: jailTopCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var unpauseCommand = &cobra.Command{
	Use:   "unpause JAIL [JAIL...]",
	Short: jailUnpauseCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var updateCommand = &cobra.Command{
	Use:   "update [OPTIONS] JAIL [JAIL...]",
	Short: jailUpdateCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var waitCommand = &cobra.Command{
	Use:   "wait JAIL [JAIL...]",
	Short: jailWaitCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

// ── image shorthands ─────────────────────────────────────────────────────────

var buildCommand = &cobra.Command{
	Use:   "build [OPTIONS] PATH | URL | -",
	Short: imageBuildCommand.Short,
	Args:  cobra.MaximumNArgs(1),
	RunE:  notImplemented,
}

var historyCommand = &cobra.Command{
	Use:   "history [OPTIONS] IMAGE",
	Short: imageHistoryCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var imagesCommand = &cobra.Command{
	Use:   "images [OPTIONS] [REPOSITORY[:TAG]]",
	Short: imageListCommand.Short,
	RunE:  notImplemented,
}

var importCommand = &cobra.Command{
	Use:   "import [OPTIONS] file|URL|- [REPOSITORY[:TAG]]",
	Short: imageImportCommand.Short,
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var loadCommand = &cobra.Command{
	Use:   "load [OPTIONS]",
	Short: imageLoadCommand.Short,
	RunE:  notImplemented,
}

var pullCommand = &cobra.Command{
	Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	Short: imagePullCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var pushCommand = &cobra.Command{
	Use:   "push [OPTIONS] NAME[:TAG]",
	Short: imagePushCommand.Short,
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var removeImageCommand = &cobra.Command{
	Use:   "rmi [OPTIONS] IMAGE [IMAGE...]",
	Short: imageRemoveCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var saveCommand = &cobra.Command{
	Use:   "save [OPTIONS] IMAGE [IMAGE...]",
	Short: imageSaveCommand.Short,
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var tagCommand = &cobra.Command{
	Use:   "tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]",
	Short: imageTagCommand.Short,
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

// ── system shorthands ────────────────────────────────────────────────────────

var infoCommand = &cobra.Command{
	Use:   "info [OPTIONS]",
	Short: systemInfoCommand.Short,
	RunE:  notImplemented,
}

var inspectCommand = &cobra.Command{
	Use:   "inspect [OPTIONS] NAME|ID [NAME|ID...]",
	Short: "Return low-level information on a jail, image, volume or network",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var searchCommand = &cobra.Command{
	Use:   "search [OPTIONS] TERM",
	Short: "Search the registry for images",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var loginCommand = &cobra.Command{
	Use:   "login [OPTIONS] [SERVER]",
	Short: "Authenticate to a registry",
	Args:  cobra.MaximumNArgs(1),
	RunE:  notImplemented,
}

var logoutCommand = &cobra.Command{
	Use:   "logout [SERVER]",
	Short: "Log out from a registry",
	Args:  cobra.MaximumNArgs(1),
	RunE:  notImplemented,
}

func init() {
	// jail shorthands — reuse shared flag helpers
	addJailRunFlags(runCommand)
	addJailRunFlags(createCommand)
	addJailListFlags(psCommand)
	addJailLogsFlags(logsCommand)
	addJailExecFlags(execCommand)
	addJailStopFlags(stopCommand)
	addJailStopFlags(restartCommand)
	addJailRemoveFlags(removeCommand)

	attachCommand.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	attachCommand.Flags().Bool("no-stdin", false, "Do not attach STDIN")
	attachCommand.Flags().Bool("sig-proxy", true, "Proxy all received signals to the process")

	commitCommand.Flags().StringP("author", "a", "", "Author")
	commitCommand.Flags().StringSliceP("change", "c", nil, "Apply a Jailfile instruction to the created image")
	commitCommand.Flags().StringP("message", "m", "", "Commit message")
	commitCommand.Flags().BoolP("pause", "p", true, "Pause jail during commit")

	copyCommand.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	copyCommand.Flags().BoolP("follow-link", "L", false, "Always follow symbol link in SRC_PATH")
	copyCommand.Flags().BoolP("quiet", "q", false, "Suppress progress output during copy")

	eventsCommand.Flags().StringP("filter", "f", "", "Filter events based on conditions provided")
	eventsCommand.Flags().String("format", "", "Format output using a custom template")
	eventsCommand.Flags().String("since", "", "Show all events created since timestamp")
	eventsCommand.Flags().String("until", "", "Stream events until this timestamp")

	exportCommand.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")

	killCommand.Flags().StringP("signal", "s", "KILL", "Signal to send to the jail")

	startCommand.Flags().BoolP("attach", "a", false, "Attach STDOUT/STDERR and forward signals")
	startCommand.Flags().String("detach-keys", "", "Override the key sequence for detaching")
	startCommand.Flags().BoolP("interactive", "i", false, "Attach jail's STDIN")

	statsCommand.Flags().BoolP("all", "a", false, "Show all jails (default shows just running)")
	statsCommand.Flags().String("format", "", "Format output using a custom template")
	statsCommand.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	statsCommand.Flags().Bool("no-trunc", false, "Do not truncate output")

	updateCommand.Flags().String("cpus", "", "Number of CPUs (RCTL limit)")
	updateCommand.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution")
	updateCommand.Flags().StringP("memory", "m", "", "Memory limit")
	updateCommand.Flags().String("memory-swap", "", "Swap limit")
	updateCommand.Flags().String("pids-limit", "", "Tune process limits via RCTL")
	updateCommand.Flags().String("restart", "", "Restart policy to apply")

	// image shorthands
	addImageBuildFlags(buildCommand)
	addImageListFlags(imagesCommand)

	historyCommand.Flags().String("format", "", "Format output using a custom template")
	historyCommand.Flags().Bool("human", true, "Print sizes and dates in human readable format")
	historyCommand.Flags().Bool("no-trunc", false, "Don't truncate output")
	historyCommand.Flags().BoolP("quiet", "q", false, "Only show image IDs")

	importCommand.Flags().StringSliceP("change", "c", nil, "Apply Jailfile instruction to the created image")
	importCommand.Flags().StringP("message", "m", "", "Set commit message for imported image")
	importCommand.Flags().String("platform", "", "Set platform if image is multi-platform capable")

	loadCommand.Flags().StringP("input", "i", "", "Read from tar archive file, instead of STDIN")
	loadCommand.Flags().BoolP("quiet", "q", false, "Suppress the load output")

	pullCommand.Flags().BoolP("all-tags", "a", false, "Download all tagged images in the repository")
	pullCommand.Flags().Bool("disable-content-trust", true, "Skip image verification")
	pullCommand.Flags().String("platform", "", "Set platform if image is multi-platform capable")
	pullCommand.Flags().BoolP("quiet", "q", false, "Suppress verbose output")

	pushCommand.Flags().BoolP("all-tags", "a", false, "Push all tags of an image to the registry")
	pushCommand.Flags().Bool("disable-content-trust", true, "Skip image signing")
	pushCommand.Flags().BoolP("quiet", "q", false, "Suppress verbose output")

	removeImageCommand.Flags().BoolP("force", "f", false, "Force removal of the image")
	removeImageCommand.Flags().Bool("no-prune", false, "Do not delete untagged parents")

	saveCommand.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")

	// system shorthands
	infoCommand.Flags().StringP("format", "f", "", "Format output using a custom template")

	inspectCommand.Flags().StringP("format", "f", "", "Format output using a custom template")
	inspectCommand.Flags().BoolP("size", "s", false, "Display total file sizes if the type is jail")

	searchCommand.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	searchCommand.Flags().String("format", "", "Format output using a custom template")
	searchCommand.Flags().Int("limit", 25, "Max number of search results")
	searchCommand.Flags().Bool("no-trunc", false, "Don't truncate output")

	loginCommand.Flags().StringP("password", "p", "", "Password")
	loginCommand.Flags().Bool("password-stdin", false, "Take the password from stdin")
	loginCommand.Flags().StringP("username", "u", "", "Username")

	rootCommand.AddCommand(
		// jail shorthands
		attachCommand,
		commitCommand,
		copyCommand,
		createCommand,
		diffCommand,
		eventsCommand,
		execCommand,
		exportCommand,
		killCommand,
		logsCommand,
		pauseCommand,
		portCommand,
		psCommand,
		renameCommand,
		restartCommand,
		removeCommand,
		runCommand,
		startCommand,
		statsCommand,
		stopCommand,
		topCommand,
		unpauseCommand,
		updateCommand,
		waitCommand,
		// image shorthands
		buildCommand,
		historyCommand,
		imagesCommand,
		importCommand,
		loadCommand,
		pullCommand,
		pushCommand,
		removeImageCommand,
		saveCommand,
		tagCommand,
		// system shorthands
		infoCommand,
		inspectCommand,
		loginCommand,
		logoutCommand,
		searchCommand,
	)
}

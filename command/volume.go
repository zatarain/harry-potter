package command

import "github.com/spf13/cobra"

var volumeCommand = &cobra.Command{
	Use:     "volume",
	Short:   "Manage volumes",
	Aliases: []string{"vol"},
}

// ── subcommands ──────────────────────────────────────────────────────────────

var volumeCreateCommand = &cobra.Command{
	Use:   "create [OPTIONS] [VOLUME]",
	Short: "Create a volume",
	Args:  cobra.MaximumNArgs(1),
	RunE:  notImplemented,
}

var volumeInspectCommand = &cobra.Command{
	Use:   "inspect [OPTIONS] VOLUME [VOLUME...]",
	Short: "Display detailed information on one or more volumes",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var volumeListCommand = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Aliases: []string{"list"},
	Short:   "List volumes",
	RunE:    notImplemented,
}

var volumePruneCommand = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove unused local volumes",
	RunE:  notImplemented,
}

var volumeRemoveCommand = &cobra.Command{
	Use:     "rm [OPTIONS] VOLUME [VOLUME...]",
	Aliases: []string{"remove"},
	Short:   "Remove one or more volumes",
	Args:    cobra.MinimumNArgs(1),
	RunE:    notImplemented,
}

func init() {
	volumeCreateCommand.Flags().StringP("driver", "d", "local", "Specify volume driver name")
	volumeCreateCommand.Flags().StringSlice("label", nil, "Set metadata for a volume")
	volumeCreateCommand.Flags().StringToStringP("opt", "o", nil, "Set driver specific options")

	volumeInspectCommand.Flags().StringP("format", "f", "", "Format output using a custom template")

	volumeListCommand.Flags().String("cluster", "", "Display only cluster volumes and use cluster volume list formatting")
	volumeListCommand.Flags().StringP("filter", "f", "", "Provide filter values (e.g. 'dangling=true')")
	volumeListCommand.Flags().String("format", "", "Format output using a custom template")
	volumeListCommand.Flags().BoolP("quiet", "q", false, "Only display volume names")

	volumePruneCommand.Flags().BoolP("all", "a", false, "Remove all unused volumes, not just anonymous ones")
	volumePruneCommand.Flags().String("filter", "", "Provide filter values (e.g. 'label=<label>')")
	volumePruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")

	volumeRemoveCommand.Flags().BoolP("force", "f", false, "Force the removal of one or more volumes")

	volumeCommand.AddCommand(
		volumeCreateCommand,
		volumeInspectCommand,
		volumeListCommand,
		volumePruneCommand,
		volumeRemoveCommand,
	)
	rootCommand.AddCommand(volumeCommand)
}

package command

import "github.com/spf13/cobra"

var systemCommand = &cobra.Command{
	Use:   "system",
	Short: "Manage harry-potter",
}

// ── subcommands ──────────────────────────────────────────────────────────────

var systemDiskUsageCommand = &cobra.Command{
	Use:   "df [OPTIONS]",
	Short: "Show harry-potter disk usage",
	RunE:  notImplemented,
}

var systemEventsCommand = &cobra.Command{
	Use:     "events [OPTIONS]",
	Aliases: []string{"ev"},
	Short:   "Get real time events from the server",
	RunE:    notImplemented,
}

var systemInfoCommand = &cobra.Command{
	Use:   "info [OPTIONS]",
	Short: "Display system-wide information",
	RunE:  notImplemented,
}

var systemPruneCommand = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove unused data",
	RunE:  notImplemented,
}

func init() {
	systemDiskUsageCommand.Flags().String("format", "", "Format output using a custom template")
	systemDiskUsageCommand.Flags().BoolP("verbose", "v", false, "Show detailed information on space usage")

	systemEventsCommand.Flags().StringP("filter", "f", "", "Filter events based on conditions provided")
	systemEventsCommand.Flags().String("format", "", "Format output using a custom template")
	systemEventsCommand.Flags().String("since", "", "Show all events created since timestamp")
	systemEventsCommand.Flags().String("until", "", "Stream events until this timestamp")

	systemInfoCommand.Flags().StringP("format", "f", "", "Format output using a custom template")

	systemPruneCommand.Flags().BoolP("all", "a", false, "Remove all unused images not just dangling ones")
	systemPruneCommand.Flags().String("filter", "", "Provide filter values (e.g. 'label=<key>=<value>')")
	systemPruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	systemPruneCommand.Flags().Bool("volumes", false, "Prune anonymous volumes")

	systemCommand.AddCommand(
		systemDiskUsageCommand,
		systemEventsCommand,
		systemInfoCommand,
		systemPruneCommand,
	)
	rootCommand.AddCommand(systemCommand)
}

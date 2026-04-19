package command

import "github.com/spf13/cobra"

var networkCommand = &cobra.Command{
	Use:     "network",
	Short:   "Manage networks",
	Aliases: []string{"net"},
}

// ── subcommands ──────────────────────────────────────────────────────────────

var networkConnectCommand = &cobra.Command{
	Use:   "connect [OPTIONS] NETWORK JAIL",
	Short: "Connect a jail to a network",
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var networkCreateCommand = &cobra.Command{
	Use:   "create [OPTIONS] NETWORK",
	Short: "Create a network",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var networkDisconnectCommand = &cobra.Command{
	Use:   "disconnect [OPTIONS] NETWORK JAIL",
	Short: "Disconnect a jail from a network",
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

var networkInspectCommand = &cobra.Command{
	Use:   "inspect [OPTIONS] NETWORK [NETWORK...]",
	Short: "Display detailed information on one or more networks",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var networkListCommand = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Aliases: []string{"list"},
	Short:   "List networks",
	RunE:    notImplemented,
}

var networkPruneCommand = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove all unused networks",
	RunE:  notImplemented,
}

var networkRemoveCommand = &cobra.Command{
	Use:     "rm NETWORK [NETWORK...]",
	Aliases: []string{"remove"},
	Short:   "Remove one or more networks",
	Args:    cobra.MinimumNArgs(1),
	RunE:    notImplemented,
}

// ── flag helpers (shared with top-level shorthands) ──────────────────────────

func addNetworkListFlags(target *cobra.Command) {
	target.Flags().StringP("filter", "f", "", "Provide filter values (e.g. 'driver=bridge')")
	target.Flags().String("format", "", "Format output using a custom template")
	target.Flags().Bool("no-trunc", false, "Do not truncate the output")
	target.Flags().BoolP("quiet", "q", false, "Only display network IDs")
}

func init() {
	networkConnectCommand.Flags().StringSlice("alias", nil, "Add network-scoped alias for the jail")
	networkConnectCommand.Flags().String("ip", "", "IPv4 address")
	networkConnectCommand.Flags().String("ip6", "", "IPv6 address")
	networkConnectCommand.Flags().StringSlice("link", nil, "Add link to another jail")
	networkConnectCommand.Flags().StringSlice("link-local-ip", nil, "Add a link-local address for the jail")

	networkCreateCommand.Flags().StringP("driver", "d", "bridge", "Driver to manage the Network (bridge)")
	networkCreateCommand.Flags().StringSlice("gateway", nil, "IPv4 or IPv6 gateway for the master subnet")
	networkCreateCommand.Flags().Bool("ingress", false, "Create swarm routing-mesh network")
	networkCreateCommand.Flags().Bool("internal", false, "Restrict external access to the network")
	networkCreateCommand.Flags().StringSlice("ip-range", nil, "Allocate jail IP from a sub-range")
	networkCreateCommand.Flags().StringToString("ipam-opt", nil, "Set IPAM driver specific options")
	networkCreateCommand.Flags().Bool("ipv6", false, "Enable IPv6 networking")
	networkCreateCommand.Flags().StringSlice("label", nil, "Set metadata on a network")
	networkCreateCommand.Flags().StringToString("opt", nil, "Set driver specific options")
	networkCreateCommand.Flags().String("scope", "", "Control the network's scope")
	networkCreateCommand.Flags().StringSlice("subnet", nil, "Subnet in CIDR format (e.g. 10.0.0.0/24)")

	networkDisconnectCommand.Flags().BoolP("force", "f", false, "Force the jail to disconnect from a network")

	networkInspectCommand.Flags().StringP("format", "f", "", "Format output using a custom template")
	networkInspectCommand.Flags().Bool("verbose", false, "Show detailed information about network resources")

	addNetworkListFlags(networkListCommand)

	networkPruneCommand.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")
	networkPruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")

	networkCommand.AddCommand(
		networkConnectCommand,
		networkCreateCommand,
		networkDisconnectCommand,
		networkInspectCommand,
		networkListCommand,
		networkPruneCommand,
		networkRemoveCommand,
	)
	rootCommand.AddCommand(networkCommand)
}

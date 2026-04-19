package command

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

var versionCommand = &cobra.Command{
	Use:   "version [OPTIONS]",
	Short: "Show the harry-potter version information",
	Run: func(invoked *cobra.Command, _ []string) {
		short, _ := invoked.Flags().GetBool("short")
		if short {
			fmt.Println(Version)
			return
		}
		fmt.Printf("harry-potter:\n")
		fmt.Printf(" Version:      %s\n", Version)
		fmt.Printf(" Go version:   %s\n", runtime.Version())
		fmt.Printf(" Git commit:   %s\n", Commit)
		fmt.Printf(" Built:        %s\n", BuildDate)
		fmt.Printf(" OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	versionCommand.Flags().Bool("short", false, "Print just the version number")
	rootCommand.AddCommand(versionCommand)
}

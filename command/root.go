package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile  string
	rootCommand = &cobra.Command{
		Use:   binaryName(),
		Short: "harry-potter — FreeBSD jail lifecycle manager",
		Long: `harry-potter is a command-line tool for managing FreeBSD jails.

It provides a unified interface for creating, configuring, and operating
jails, virtual networks, ZFS-backed snapshots, and base system templates.`,
		SilenceUsage: true,
	}
)

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVar(&configFile, "config", "", "config file (default: $HOME/.harry-potter/config.yaml)")
	rootCommand.PersistentFlags().Bool("verbose", false, "enable verbose output")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, ".harry-potter"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}

func binaryName() string {
	return filepath.Base(os.Args[0])
}

func notImplemented(invoked *cobra.Command, _ []string) error {
	return fmt.Errorf("%q is not yet implemented", invoked.CommandPath())
}

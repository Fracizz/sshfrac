package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgPath  string
	insecure bool
	// Version is overwritten by -ldflags at build time.
	Version = "0.2.0"
)

var rootCmd = &cobra.Command{
	Use:           "sshctl",
	Short:         "High-performance SSH/SCP CLI with encrypted server inventory",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       Version,
}

// Execute runs the root command.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to servers JSON (default: ~/.sshctl/servers.json or $SSHCTL_CONFIG)")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "skip SSH host key verification (unsafe; for lab only)")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(scpCmd)
	rootCmd.AddCommand(initCmd)
}

package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Stores the remote registry in the configuration directory",
	Long: `Stores the remote registry in the configuration directory.

It always override the existing (local) registry with the newest remote one.`,
	Run: func(cmd *cobra.Command, args []string) {
		registryUpgrade(env.DefaultConfigEnv())
	},
}

func registryUpgrade(e *env.ConfigEnv) {
	err := registry.Upgrade(e)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to upgrade registry: %v\n", err)
		nonZeroExit(1)
	}
}

func init() {
	registryCmd.AddCommand(upgradeCmd)
}

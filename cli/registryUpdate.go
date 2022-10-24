package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Compares local and remote registry",
	Long: `Compares local and remote registry.

In the case that there are differences, it prints the diff on the console.
If there aren't any, it exits successfully.`,
	Run: func(cmd *cobra.Command, args []string) {
		isJson, _ := strconv.ParseBool(rootCmd.Flag("json").Value.String())
		registryUpdate(env.DefaultConfigEnv(isJson))
	},
}

func registryUpdate(e *env.ConfigEnv) {
	s, err := registry.UpdateRegistry(e)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to update registry: %v\n", err)
		nonZeroExit(1)
	}

	_, _ = fmt.Fprintf(e.Writer(), s)
}
func init() {
	registryCmd.AddCommand(updateCmd)
}

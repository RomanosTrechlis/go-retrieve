package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Delete configuration folder with the corresponding files",
	Long:  `Delete configuration folder with the corresponding files`,
	Run: func(cmd *cobra.Command, args []string) {
		executeDestroy(env.DefaultConfigEnv())
	},
}

func executeDestroy(e *env.ConfigEnv) {
	err := config.Destroy(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to destroy environment: %v\n", err)
		nonZeroExit(1)
	}
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}

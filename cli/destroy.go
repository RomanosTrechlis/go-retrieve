package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"github.com/spf13/cobra"
	"os"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Delete configuration folder with the corresponding files",
	Long:  `Delete configuration folder with the corresponding files`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Destroy()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to destroy environment: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}

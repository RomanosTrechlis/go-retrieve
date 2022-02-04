package cli

import (
	"github.com/spf13/cobra"
)

// registryCmd represents the registry command
var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Not yet implemented",
	Long:  `Not yet implemented`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("registry not yet implemented")
	//},
}

func init() {
	rootCmd.AddCommand(registryCmd)
}

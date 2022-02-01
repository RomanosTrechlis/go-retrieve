package cli

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"github.com/RomanosTrechlis/retemp/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display the configuration of specific profile",
	Long:  `Display the configuration of specific profile`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "provide at least one profile to inspect\n")
			os.Exit(1)
		}

		c, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}

		isFirst := true
		for _, p := range c.Profiles {
			if !util.Contains(args, p.Name) {
				continue
			}

			if !isFirst {
				isFirst = false
				fmt.Printf("------------- %s -----------\n", p.Name)
			}

			if d, _ := cmd.Flags().GetBool("dump"); d {
				spew.Dump(p)
				os.Exit(0)
			}
			s, _ := json.MarshalIndent(p, "", "  ")
			fmt.Println(string(s))
		}
	},
}

func init() {
	profileCmd.AddCommand(inspectCmd)
	inspectCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
}

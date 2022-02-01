package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"os"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Applies changes to profile",
	Long: `Display active profile, display all available profiles and change active profile.

Additionally, you can inspect a specific profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}

		if len(args) > 1 {
			_, _ = fmt.Fprintf(os.Stderr, "too many arguments received: %v", args)
			os.Exit(1)
		}

		if len(args) == 1 {
			err = updateActive(c, args[0])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		if c.Active == nil && len(c.Profiles) == 1 {
			err = updateActive(c, c.Profiles[0].Name)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
				os.Exit(1)
			}
		}

		displayActive(c)
	},
}

func updateActive(c *config.Configuration, profile string) error {
	for _, p := range c.Profiles {
		if p.Name == profile {
			c.Active = p
			return config.UpdateConfig(c)
		}
	}
	return fmt.Errorf("failed to find profile '%s' in the configuration", profile)
}

func displayActive(c *config.Configuration) {
	profile := "NOT SET"
	if c.Active != nil {
		profile = c.Active.Name
	}

	fmt.Printf("active profile: %s", profile)
}

func init() {
	rootCmd.AddCommand(profileCmd)
}

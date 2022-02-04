package cli

import (
	"fmt"
	"os"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Applies changes to profile",
	Long: `Display active profile, display all available profiles and change active profile.

Additionally, you can inspect a specific profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		executeProfile(env.DefaultConfigEnv(), args)
	},
}

func executeProfile(e *env.ConfigEnv, args []string) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		nonZeroExit(1)
	}

	if len(args) > 1 {
		_, _ = fmt.Fprintf(os.Stderr, "too many arguments received: %v", args)
		nonZeroExit(1)
	}

	if len(args) == 1 {
		err = updateActive(e, c, args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
			nonZeroExit(1)
		}
		return
	}

	if c.Active == nil && len(c.Profiles) == 1 {
		err = updateActive(e, c, c.Profiles[0].Name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
			nonZeroExit(1)
		}
	}

	displayActive(e, c)
}

func updateActive(e *env.ConfigEnv, c *config.Configuration, profile string) error {
	for _, p := range c.Profiles {
		if p.Name == profile {
			c.Active = p
			return config.UpdateConfig(e, c)
		}
	}
	return fmt.Errorf("failed to find profile '%s' in the configuration", profile)
}

func displayActive(e *env.ConfigEnv, c *config.Configuration) {
	profile := "NOT SET"
	if c.Active != nil {
		profile = c.Active.Name
	}

	_, _ = fmt.Fprintf(e.Writer(), "active profile: %s", profile)
}

func init() {
	rootCmd.AddCommand(profileCmd)
}

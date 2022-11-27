package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

func RegisterProfileCmd(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "profile",
		Short: "Applies changes to profile",
		Long: `Display active profile, display all available profiles and change active profile.

Additionally, you can inspect a specific profile.`,
		Run: func(cmd *cobra.Command, args []string) {
			executeProfile(env.DefaultConfigEnv(false), args)
		},
	}
}

func RegisterProfileListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all the available profiles",
		Long:  `Lists all the available profiles`,
		Run: func(cmd *cobra.Command, args []string) {
			executeProfileList(env.DefaultConfigEnv(false))
		},
	}
}

func RegisterProfileInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect",
		Short: "Display the configuration of specific profile",
		Long:  `Display the configuration of specific profile`,
		Run: func(cmd *cobra.Command, args []string) {
			dump, _ := cmd.Flags().GetBool("dump")
			executeProfileInspect(env.DefaultConfigEnv(false), args, dump, false)
		},
	}
}

func executeProfileInspect(e *env.ConfigEnv, args []string, dump, isJson bool) {
	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "provide at least one profile to inspect\n")
		NonZeroExit(1)
	}

	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		NonZeroExit(1)
	}

	isFirst := true
	for _, p := range c.Profiles {
		if !util.Contains(args, p.Name) {
			continue
		}

		if !isFirst {
			isFirst = false
			_, _ = fmt.Fprintf(e.Writer(), "------------- %s -----------\n", p.Name)
		}

		if dump {
			spew.Fdump(e.Writer(), p)
			return
		}
		s, _ := util.MarshalIndent(c, isJson)
		_, _ = fmt.Fprintln(e.Writer(), string(s))
	}
}

func executeProfileList(e *env.ConfigEnv) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		NonZeroExit(1)
	}

	for _, p := range c.Profiles {
		_, _ = fmt.Fprintln(e.Writer(), p.Name)
	}
}

func executeProfile(e *env.ConfigEnv, args []string) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		NonZeroExit(1)
	}

	if len(args) > 1 {
		_, _ = fmt.Fprintf(os.Stderr, "too many arguments received: %v", args)
		NonZeroExit(1)
	}

	if len(args) == 1 {
		err = updateActive(e, c, args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
			NonZeroExit(1)
		}
		return
	}

	if c.Active == nil && len(c.Profiles) == 1 {
		err = updateActive(e, c, c.Profiles[0].Name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to update active profile: %v\n", err)
			NonZeroExit(1)
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

	_, _ = fmt.Fprintf(e.Writer(), "active profile: %s\n", profile)
}

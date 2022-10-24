package main

import "github.com/RomanosTrechlis/go-retrieve/cli"

func main() {
	// init: create config file holding profiles and sources
	// destroy: delete config directory and files

	// config reinit (?)
	// config: prints configuration file

	// profile: show active profile
	// profile <profile name>: activates the profile name specified
	// profile list: show available profiles
	// profile inspect <profile name>: show profile's sources

	// registry define: wizard to create registry file for source
	// registry update: returns diffs between local registry and remote registry
	// registry upgrade: downloads new versions of template definition from remote sources

	// template <template name>: find first instance of template name and download the corresponding directory
	// template list: lists all templates
	// template list <template name>: displays tree of files of the corresponding template
	// template create <directory>: takes a directory zips it, asks name, after unzip commands etc

	// backup: create a backup file of the configuration file

	cli.Execute()
}

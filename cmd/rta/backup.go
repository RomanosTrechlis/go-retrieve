package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// backupCmd represents the backup command
var backupCmd = cli.RegisterBackupCmd()

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringP("filename", "f", "", "Backup to specific file")
}

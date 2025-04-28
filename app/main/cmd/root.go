/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "main is a CLI application for show golang commands",
	Long:  `This application is a tool to show golang commands.`,
}

/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagUsername = "username"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello is a command to print hello",
	Run: func(_ *cobra.Command, _ []string) {
		username := viper.GetString(fmt.Sprintf("hello.%s", FlagUsername))
		fmt.Printf("Hello, %s!\n", username)
	},
}

func init() {
	helloCmd.Flags().String(FlagUsername, "me", "say hello to this user")
	_ = viper.BindPFlag(fmt.Sprintf("hello.%s", FlagUsername), helloCmd.Flags().Lookup(FlagUsername))
	rootCmd.AddCommand(helloCmd)
}

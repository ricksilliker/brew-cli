package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shellCmd)
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a terminal with a environment.",
	Long:  "Start a terminal with a environment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Started shell.")
	},
}

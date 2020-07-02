package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Get the latest project configurations.",
	Long:  "Get the latest project configurations.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pulled Latest!")
	},
}
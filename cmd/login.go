package cmd

import (
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var loginCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get project info.",
	Long:  "Get Brazen Animation project information from Shotgun.",
	Run: func(cmd *cobra.Command, args []string) {
		login(cmd)
	},
}

func init() {
	loginCmd.Flags().Bool("json", false, "Output in json blobs.")
	loginCmd.Flags().String("user", "", "Shotgun user name.")
	loginCmd.Flags().String("pass", "", "Shotgun user password.")

	rootCmd.AddCommand(loginCmd)
}

func login(cmd *cobra.Command) {
	asJson, _ := cmd.Flags().GetBool("json")
	user, _ := cmd.Flags().GetString("user")
	pass, _ := cmd.Flags().GetString("pass")

	if asJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	_, err := brew.AuthenticateAsUser(user, pass)
	if err != nil {
		logrus.WithError(err).Error("failed to authenticate user")
		os.Exit(1)
		return
	}
	os.Exit(0)
}
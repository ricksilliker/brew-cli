package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List available software.",
	Long:  "Get information about available applications that can be run with this tool.",
	Run: func(cmd *cobra.Command, args []string) {
		apps(cmd)
	},
}

func init() {
	appsCmd.Flags().Bool("json", false, "Output in json blobs.")
	appsCmd.Flags().String("user", "", "Shotgun user name.")
	appsCmd.Flags().String("pass", "", "Shotgun user password.")

	rootCmd.AddCommand(appsCmd)
}

func apps(cmd *cobra.Command) {
	asJson, _ := cmd.Flags().GetBool("json")
	user, _ := cmd.Flags().GetString("user")
	pass, _ := cmd.Flags().GetString("pass")

	if asJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	a, err := brew.AuthenticateAsUser(user, pass)
	if err != nil {
		logrus.WithError(err).Error("failed to authenticate user")
		os.Exit(1)
		return
	}
	token := fmt.Sprintf("%v %v", a.TokenType, a.AccessToken)
	sf := brew.GetAllSoftware(token, user)
	enc := json.NewEncoder(os.Stdout)
	d := map[string][]brew.Software{"data": sf}
	_ = enc.Encode(d)
	os.Exit(0)
}
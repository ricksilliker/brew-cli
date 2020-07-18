package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get project info.",
	Long:  "Get Brazen Animation project information from Shotgun.",
	Run: func(cmd *cobra.Command, args []string) {
		resolveProjectsQuery(cmd)
	},
}

func init() {
	projectsCmd.Flags().Bool("json", false, "Output in json blobs.")
	projectsCmd.Flags().String("user", "", "Shotgun user name.")
	projectsCmd.Flags().String("pass", "", "Shotgun user password.")

	rootCmd.AddCommand(projectsCmd)
}

func resolveProjectsQuery(cmd *cobra.Command) {
	asJson, _ := cmd.Flags().GetBool("json")
	user, _ := cmd.Flags().GetString("user")
	pass, _ := cmd.Flags().GetString("pass")

	if asJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	a := brew.AuthenticateAsUser(user, pass)
	token := fmt.Sprintf("%v %v", a.TokenType, a.AccessToken)
	proj := brew.GetAllProjects(token, user)
	enc := json.NewEncoder(os.Stdout)
	d := map[string][]brew.Project{"data": proj}
	_ = enc.Encode(d)
	os.Exit(0)
}
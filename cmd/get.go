package cmd

import (
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GetOpts struct {
	Args BrazenOpts
	Environment bool
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get project information.",
	Long:  "Get eco data related to a specific project like the environment, available apps, and assigned projects.",
	Run: func(cmd *cobra.Command, args []string) {
		brazenOpts := ParseGlobalFlags(cmd.Flags())
		if brazenOpts.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		env, err := cmd.Flags().GetBool("environment")
		if err != nil {
			logrus.Fatal(err)
		}

		getData(&GetOpts{
			Args: brazenOpts,
			Environment: env,
		})
	},
}

func init() {
	getCmd.Flags().Bool("environment", false, "Print a JSON blob of the eco environment.")
	rootCmd.AddCommand(getCmd)
}

func getData(opts *GetOpts) {
	ctx := brew.BrewContext{
		Project: opts.Args.Eco,
		EcoDirectory: opts.Args.EcoDir,
	}

	rootEcoFile := brew.ResolveEco(&ctx)
	rawEnvironment := brew.GetRawEnvironment(rootEcoFile)
	logrus.Info(rawEnvironment)

	if opts.Environment {
		logrus.Info("Handle get environment command.")
	}
}
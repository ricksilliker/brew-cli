package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "brazen",
	Short: "Brazen Animation environment wrapper.",
	Long: "Use this to create, manage, and launch a project.",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Base command ran.")
	},
}

var defaultTools []string

func init() {
	rootCmd.PersistentFlags().String("project", "", "Project name or code.")
	rootCmd.PersistentFlags().String("shot", "", "Shot name/code.")
	rootCmd.PersistentFlags().String("bundle", "", "Application environment context name.")
	rootCmd.PersistentFlags().StringArray("tools", defaultTools, "Comma separated list of tools.")
	rootCmd.PersistentFlags().String("eco-dir", "eco", "Directory path where the eco file is located.")
	rootCmd.PersistentFlags().Bool("debug", false, "Make output more verbose.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ParseFlags(flags *pflag.FlagSet) brew.BrewContext {
	project, _ := flags.GetString("project")
	shot, _ := flags.GetString("shot")
	bundle, _ := flags.GetString("bundle")
	tools, _ := flags.GetStringArray("tools")
	eco, _ := flags.GetString("eco-dir")

	cwd, err := os.Getwd()
	if err != nil {
		println(err)
	}

	env := brew.EnvironmentContext{}

	return brew.BrewContext{
		Project: project,
		Shot: shot,
		Bundle: bundle,
		Tools: tools,
		Eco: eco,
		CurrentDirectory: cwd,
		Site: "dallas",
		Environment: env,
	}
}
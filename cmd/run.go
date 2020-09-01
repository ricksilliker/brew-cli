package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
)

type RunOpts struct {
	Args BrazenOpts
	CommandName string
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use an application with a specific environment.",
	Long:  "Use an application with a specific environment.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logrus.Error("No command given to run, exiting.")
		}

		brazenOpts := ParseGlobalFlags(cmd.Flags())

		runApplication(&RunOpts{
			Args: brazenOpts,
			CommandName: args[0],
		})
	},
	Args: cobra.MaximumNArgs(1),

}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runApplication(opts *RunOpts) {
	ctx := brew.BrewContext{
		Project: opts.Args.Eco,
		EcoDirectory: opts.Args.EcoDir,
	}

	eco := brew.ResolveEco(&ctx)
	env := brew.GetRawEnvironment(eco)
	var serializedEnv []string
	for key, value := range env {
		serializedValue := fmt.Sprintf("%v=%v", key, value)
		serializedEnv = append(serializedEnv, serializedValue)
	}

	proc := exec.Command(opts.CommandName)
	proc.Env = serializedEnv
	err := proc.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
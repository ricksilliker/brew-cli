package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/spf13/cobra"
	"os/exec"
)

type RunContext struct {
	Site string
	Eco string
	Project string
	ToolRequests []string
	Bundle string
	Shot string
}

var defaultTools []string
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use an application with a specific environment.",
	Long:  "Use an application with a specific environment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		if len(args) == 0 {
			fmt.Println("No command given to run, exiting.")
			return
		}
		fmt.Println("Ran an app.")
		runApplication(cmd, args[0])
	},
	Args: cobra.MaximumNArgs(1),

}

func init() {
	runCmd.Flags().String("project", "", "Project code.")
	runCmd.Flags().String("shot", "", "Shot code.")
	runCmd.Flags().String("bundle", "", "Application environment context name.")
	runCmd.Flags().StringArray("tools", defaultTools, "Comma separated list of tools.")

	rootCmd.AddCommand(runCmd)
}

func runApplication(cmd *cobra.Command, app string) {
	project, _ := cmd.Flags().GetString("project")
	shot, _ := cmd.Flags().GetString("shot")
	bundle, _ := cmd.Flags().GetString("bundle")
	tools, _ := cmd.Flags().GetStringArray("tools")

	rootContext := ParseGlobalFlags(cmd.Flags())

	ctx := brew.BrewContext{
		Site:         rootContext.Site,
		Eco:          rootContext.EcoDir,
		Project:      project,
		Tools:        tools,
		Bundle:       bundle,
		Shot:         shot,
	}

	contextEnv := brew.GetEnv(&ctx)
	var serializedEnv []string
	for key, value := range contextEnv {
		serializedValue := fmt.Sprintf("%v=%v", key, value)
		serializedEnv = append(serializedEnv, serializedValue)
	}
	
	proc := exec.Command(app)
	proc.Env = serializedEnv
	err := proc.Run()
	if err != nil{
		fmt.Println(err)
	}
}
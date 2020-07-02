package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

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
		ctx := ParseFlags(cmd.PersistentFlags())
		run_app(ctx, args[0])
	},
	Args: cobra.MaximumNArgs(1),

}

//func init() {
//	runCmd.Flags().String("command")
//}

func run_app(ctx *brew.BrewContext, cmd string) {
	proc := exec.Command(cmd)
	proc.Env = brew.GetEnv(ctx.Environment)
	err := proc.Run()
	if err != nil{
		fmt.Println(err)
	}
}
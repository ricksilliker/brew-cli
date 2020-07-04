package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

type RootContext struct {
	Site string
	EcoDir string
	Debug bool
}

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "brazen",
	Short: "Brazen Animation environment wrapper.",
	Long: "Use this to create, manage, and launch a project.",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("Brazen Animation CLI version %v", version)
	},
}

func init() {
	rootCmd.PersistentFlags().String("site", "dallas", "Studio site location.")
	rootCmd.PersistentFlags().String("eco-dir", "", "Directory path where the eco file is located.")
	rootCmd.PersistentFlags().Bool("debug", false, "Make output more verbose.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ParseGlobalFlags(flags *pflag.FlagSet) RootContext {
	site, _ := flags.GetString("site")
	eco, _ := flags.GetString("eco-dir")
	debug, _ := flags.GetBool("debug")

	var ecoDir string
	if eco == "" {
		ecoDir = brew.GetEcoDirectory()
	} else {
		ecoDir = eco
	}

	return RootContext{
		Site:   site,
		EcoDir: ecoDir,
		Debug:  debug,
	}
}
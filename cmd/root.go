package cmd

import (
	"fmt"
	"github.com/ricksilliker/brew-cli/brew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

type BrazenOpts struct {
	Eco string
	EcoDir string
	Debug bool
}

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "brazen",
	Short: "Brazen Animation environment wrapper.",
	Long: "Use this to create, manage, and launch a project.",
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := cmd.Flags().GetBool("version")
		if err != nil {
			logrus.Fatal(err)
		}
		if ver {
			fmt.Printf("v%v", version)
			return
		}
	},
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rootCmd.Flags().Bool("version", false, "Print out CLI version.")

	rootCmd.PersistentFlags().String("eco", "", "Name of the eco file to use.")
	rootCmd.PersistentFlags().String("ecoDirectory", "", "Directory where eco files are located.")
	rootCmd.PersistentFlags().Bool("debug", false, "Include debug logs in output.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ParseGlobalFlags(flags *pflag.FlagSet) BrazenOpts {
	eco, err := flags.GetString("eco")
	if err != nil {
		logrus.Fatal(err)
	}

	ecoDir, err := flags.GetString("ecoDirectory")
	if err != nil {
		logrus.Fatal(err)
	}
	if ecoDir == "" {
		ecoDir = brew.GetEcoDirectory()
	}

	debug, err := flags.GetBool("debug")
	if err != nil {
		logrus.Fatal(err)
	}

	return BrazenOpts{
		Eco:   eco,
		EcoDir: ecoDir,
		Debug:  debug,
	}
}
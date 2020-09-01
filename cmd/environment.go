package cmd

import (
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Get the relevant environment keys/values.",
	Long:  "Get environment variables relevant to a given project, bundle or both.",
	Run: func(cmd *cobra.Command, args []string) {
		resolveEnvironmentQuery(cmd)
	},
}

func init() {
	environmentCmd.Flags().Bool("json", false, "Output in json blobs.")
	environmentCmd.Flags().String("project", "", "Project code.")
	environmentCmd.Flags().String("bundle", "", "Application environment context name.")

	rootCmd.AddCommand(environmentCmd)
}

func resolveEnvironmentQuery(cmd *cobra.Command) {
	//asJson, _ := cmd.Flags().GetBool("json")
	//project, _ := cmd.Flags().GetString("project")
	//bundle, _ := cmd.Flags().GetString("bundle")
	//
	//if asJson {
	//	logrus.SetFormatter(&logrus.JSONFormatter{})
	//}
	//
	//rootContext := ParseGlobalFlags(cmd.Flags())
	//
	//ctx := brew.BrewContext{
	//	Site:         rootContext.Site,
	//	Eco:          rootContext.EcoDir,
	//	Project:      project,
	//	Bundle:       bundle,
	//}
	//
	//contextEnv := brew.GetEnv(&ctx)
	//if asJson {
	//	data, err := json.Marshal(contextEnv)
	//	if err != nil {
	//		logrus.Error(err)
	//	}
	//	logrus.WithField("environment", string(data)).Info("Environment settings.")
	//} else {
	//	for key, value := range contextEnv {
	//		serializedValue := fmt.Sprintf("%v=%v", key, value)
	//		fmt.Println(serializedValue)
	//	}
	//
	//}
}
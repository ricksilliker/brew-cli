package brew

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

func GetEnv(ctx *BrewContext) map[string]string {
	//Get site eco
	//Get project eco
	//Get tool ecos
	//Loop through each eco
	//Check that each as an environment key
	//store each map in a master map
	//Turn master map into string array

	var result []Eco
	result = ResolveContextEcoFiles(*ctx, result)
	result = ResolveBundleEcoFiles(*ctx, result)
	result = ResolveToolEcoFiles(*ctx, result)

	// Create empty map
	// Set parent environ
	// Add parent environ to map
	var rawEnv  = map[string]string{}
	for _, eco := range result {
		ecoFile := eco.ReadEcoFile()
		for _, item := range ecoFile.Environment {
			value := os.ExpandEnv(item.Value.(string))
			if runtime.GOOS != "windows" {
				value = strings.ReplaceAll(value, ";", ":")
			}
			err := os.Setenv(item.Key.(string), value)
			if err != nil {
				logrus.Errorf("failed to set environment variable: %v", item.Key.(string))
			}
			rawEnv[item.Key.(string)] = os.Getenv(item.Key.(string))
		}
	}

	return rawEnv
}

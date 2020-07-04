package brew

import "github.com/sirupsen/logrus"

func GetEnv(ctx *BrewContext) []string {
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

	for _, r := range result {
		logrus.Info(r.Name)
	}

	return nil
}

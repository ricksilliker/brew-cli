package brew

func GetEnv(ctx *BrewContext) []Eco {
	//Get site eco
	//Get project eco
	//Get tool ecos
	//Loop through each eco
	//Check that each as an environment key
	//store each map in a master map
	//Turn master map into string array

	var result []Eco

	contextEcos := ResolveContextEcoFiles(*ctx)
	for i, _ := range(contextEcos) {
		result = append(result, contextEcos[i])
	}



	return result
}

package brew

type EnvironmentContext struct {
	Site string `json:"site"`
	Eco string `json:"eco_dir"`
	Project string `json:"project"`
	ToolRequests []string `json:"tool_requests"`
	Bundle string `json:"bundle"`
	LazyLoad bool `json:"delayed_load"`
}

func GetEnv(ctx *EnvironmentContext) []string {
	//Get site eco
	//Get project eco
	//Get tool ecos
	//Loop through each eco
	//Check that each as an environment key
	//store each map in a master map
	//Turn master map into string array

	var result []string
	result.append()
}

package brew

type BrewContext struct {
	Site string `json:"site"`
	Shot string `json:"shot"`
	Project string `json:"project"`
	Tools []string `json:"tools"`
	Bundle string `json:"bundle"`
	CurrentDirectory string `json:"cur_dir"`
	Eco string `json:"eco_dir"`
	Environment EnvironmentContext `json:"env"`
}
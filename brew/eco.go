package brew

import (
	"fmt"
	"path"
)

type Eco struct {
	Name string
	EcoDirectory string
	Type string
	Level string
}

type EcoFile struct {
	Environment map[string] string `yaml:"environment"`
	Bundles map[string] EcoFileBundle `yaml:"bundles"`
	Requires []string `yaml:"requires"`
}

type EcoFileBundle struct {
	DefaultCmd string `yaml:"default_cmd"`
	Tools []string `yaml:"tools"`
}

func ResolveContextEcoFiles(ctx BrewContext) []Eco {
	var result []Eco
	if ctx.Site != "" {
		siteEco := Eco{
			Name:         ctx.Site,
			EcoDirectory: ctx.Eco,
			Type:         "context",
			Level:        "site",
		}
		result = append(result, siteEco)
	}

	if ctx.Project != "" {
		projectEco := Eco{
			Name:         ctx.Project,
			EcoDirectory: ctx.Eco,
			Type:         "context",
			Level:        "project",
		}
		result = append(result, projectEco)
	}

	return result
}

func ResolveToolEcoFiles(ctx BrewContext) []Eco {
	var result []Eco

	if ctx.Bundle == "" {
		return nil
	}



	return result
}

func (e *Eco) ReadEcoFile() EcoFile {

	var ecoFileName string
	if e.Type != "" {
		ecoFileName = fmt.Sprintf("%v.%v.eco", e.Name, e.Type)
	} else {
		ecoFileName = fmt.Sprintf("%v.eco", e.Name)
	}

	fp := path.Join(e.EcoDirectory, ecoFileName)




	return result
}
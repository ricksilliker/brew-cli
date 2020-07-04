package brew

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Eco struct {
	Name string
	EcoDirectory string
	Type string
	Level string
}

type EcoFile struct {
	Environment yaml.MapSlice `yaml:"environment"`
	Bundles map[string]EcoFileBundle `yaml:"bundles"`
	Requires []string `yaml:"requires"`
}

type EcoFileBundle struct {
	DefaultCmd string `yaml:"default_cmd"`
	Tools []string `yaml:"tools"`
}

func ResolveContextEcoFiles(ctx BrewContext, ecoList []Eco) []Eco {
	//if ctx.Site != "" {
	siteEco := Eco{
		Name:         ctx.Site,
		EcoDirectory: ctx.Eco,
		Type:         "context",
		Level:        "site",
	}
	ecoList = append(ecoList, siteEco)
	//}

	//if ctx.Project != "" {
	projectEco := Eco{
		Name:         ctx.Project,
		EcoDirectory: ctx.Eco,
		Type:         "context",
		Level:        "project",
	}
	ecoList = append(ecoList, projectEco)
	//}

	return ecoList
}

func ResolveBundleEcoFiles(ctx BrewContext, ecoList []Eco) []Eco {
	if ctx.Bundle == "" {
		return ecoList
	}

	var tools []string
	for _, e := range ecoList {
		if e.Type != "context" {
			continue
		}
		t := e.GetBundleTools(ctx.Bundle)
		if len(t) > 0 {
			tools = append(tools, t...)
		}
	}

	for _, t := range tools {
		toolEco := Eco{
			Name: t,
			EcoDirectory: ctx.Eco,
			Type: "tool",
		}
		ecoList = append(ecoList, toolEco)
	}

	return ecoList
}

func ResolveToolEcoFiles(ctx BrewContext, ecoList []Eco) []Eco {
	for index, _ := range ctx.Tools {
		toolEco := Eco{
			Name: ctx.Tools[index],
			EcoDirectory: ctx.Eco,
			Type: "tool",
		}
		ecoList = append(ecoList, toolEco)
	}

	return ecoList
}


func (e *Eco) GetBundleTools(bundleName string) []string {
	ecoFile := e.ReadEcoFile()
	if val, ok := ecoFile.Bundles[bundleName]; ok {
		return val.Tools
	} else {
		var empty []string
		return empty
	}
}

func (e *Eco) ReadEcoFile() *EcoFile {
	var ecoFileName string
	if e.Type != "" {
		if e.Level != "" {
			if e.Name != "" {
				ecoFileName = fmt.Sprintf("%v.%v.%v.eco", e.Name, e.Level, e.Type)
			} else {
				ecoFileName = fmt.Sprintf("%v.%v.eco", e.Level, e.Type)
			}
		} else {
			ecoFileName = fmt.Sprintf("%v.%v.eco", e.Name, e.Type)
		}
	} else {
		ecoFileName = fmt.Sprintf("%v.eco", e.Name)
	}

	fp := path.Join(e.EcoDirectory, ecoFileName)
	logrus.Info(fp)
	fileData, err := ioutil.ReadFile(fp)
	if err != nil {
		logrus.Error("failed to read eco file")
	}

	ef := EcoFile{}
	err = yaml.Unmarshal(fileData, &ef)
	if err != nil {
		logrus.Error("failed to get yaml data from eco file")
	}

	return &ef
}
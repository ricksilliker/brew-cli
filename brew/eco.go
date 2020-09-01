package brew

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)


type EcoFile struct {
	Name string
	Environment yaml.MapSlice `yaml:"environment"`
	Bundles []string `yaml:"bundles"`
	Command string `yaml:"command"`
	Inherit bool `yaml:"inherit"`
}

//type Eco struct {
//	Name string
//	EcoDirectory string
//	Type string
//	Level string
//}
//
//type EcoFile struct {
//	Environment yaml.MapSlice `yaml:"environment"`
//	Bundles map[string]EcoFileBundle `yaml:"bundles"`
//	Requires []string `yaml:"requires"`
//}
//
//type EcoFileBundle struct {
//	DefaultCmd string `yaml:"default_cmd"`
//	Tools []string `yaml:"tools"`
//}

func ResolveEco(ctx *BrewContext) *EcoFile {
	fp := path.Join(ctx.EcoDirectory, ctx.Project) + ".yaml"
	logrus.WithField("filepath", fp).Debug("Resolving eco with path")

	fileData, err := ioutil.ReadFile(fp)
	if err != nil {
		logrus.Error("failed to read eco file")
		return nil
	}
	logrus.Info(string(fileData))

	ef := EcoFile{}
	err = yaml.Unmarshal(fileData, &ef)
	if err != nil {
		logrus.Error("failed to get yaml data from eco file")
		return nil
	}
	ef.Name = ctx.Project

	return &ef
}

func GetRawEnvironment(eco *EcoFile) map[string]string {
	var rawEnv  = map[string]string{}
	rawEnv["ECO_NAME"] = eco.Name

	for _, item := range eco.Environment {
		var rawValue string
		switch d := item.Value.(type) {
		case string:
			rawValue = d
		case int:
			rawValue = strconv.Itoa(d)
		case []string:
			rawValue = strings.Join(d[:], string(os.PathListSeparator))
		default:
			continue
		}
		logrus.WithFields(logrus.Fields{
			"Key": item.Key.(string),
			"Value": rawValue,
		}).Info("Getting value from environment.")
		rawEnv[item.Key.(string)] = rawValue

		//value := os.ExpandEnv(rawValue)
		//if runtime.GOOS != "windows" {
		//	value = strings.ReplaceAll(value, ";", ":")
		//}
		//err := os.Setenv(item.Key.(string), value)
		//if err != nil {
		//	logrus.Errorf("failed to set environment variable: %v", item.Key.(string))
		//}
		//rawEnv[item.Key.(string)] = os.Getenv(item.Key.(string))
	}

	return rawEnv
}
package brew

import (
	"os"
	"runtime"
)

type BrewContext struct {
	Project string
	EcoDirectory string
}

const defaultEcoDirWindows = "O:/sww/eco"
const defaultEcoDirMacOS = "/Volumes/prod/sww/eco"
const defaultEcoDirLinux = "/mnt/prod/sww/eco"

func GetEcoDirectory() string {
	ecoDir := os.Getenv("ECO_DIR")
	if ecoDir == "" {
		if runtime.GOOS == "windows" {
			return defaultEcoDirWindows
		} else if runtime.GOOS == "darwin" {
			return defaultEcoDirMacOS
		} else {
			return defaultEcoDirLinux
		}
	}
	return ecoDir
}
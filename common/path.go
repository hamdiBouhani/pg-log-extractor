package common

import (
	"os"
)

// GetProjectPath returns the root path of project.
// ProjectPath here is the working dir of the process
func GetProjectPath() string {
	wd, _ := os.Getwd()
	return wd
}

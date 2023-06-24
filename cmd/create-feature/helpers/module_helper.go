package helpers

import (
	"os"

	"golang.org/x/mod/modfile"
)

func GetModuleName() string {
	modContent, _ := os.ReadFile("./go.mod")
	return modfile.ModulePath(modContent)
}

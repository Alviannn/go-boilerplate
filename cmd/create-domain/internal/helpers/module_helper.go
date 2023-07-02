package helpers

import "strings"

func GetModuleName() string {
	content, _ := ReadFileAsString("go.mod")
	firstLine := strings.Split(content, "\n")[0]

	return strings.Replace(firstLine, "module ", "", 1)
}

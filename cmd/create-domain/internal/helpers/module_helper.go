package helpers

import (
	"bufio"
	"os"
	"strings"
)

func GetModuleName() string {
	file, _ := os.OpenFile("go.mod", os.O_RDONLY, os.ModePerm)
	defer file.Close()

	var firstLine string

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine = scanner.Text()
	}

	return strings.Replace(firstLine, "module ", "", 1)
}

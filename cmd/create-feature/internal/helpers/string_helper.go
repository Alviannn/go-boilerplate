package helpers

import (
	"fmt"
	"strings"
)

func SnakeToPascalCase(snake string) string {
	lowerSnake := strings.ToLower(snake)
	partList := strings.Split(lowerSnake, "_")

	for i, part := range partList {
		firstChar := part[:1]
		restOfChars := part[1:]

		partList[i] = fmt.Sprintf("%s%s", strings.ToUpper(firstChar), restOfChars)
	}

	return strings.Join(partList, "")
}

func SnakeToPackageName(snake string) string {
	return strings.ReplaceAll(snake, "_", "")
}

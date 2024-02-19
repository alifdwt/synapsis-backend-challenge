package util

import "strings"

func ConvertTitleToId(title string) string {
	input := strings.ToLower(title)
	output := strings.ReplaceAll(input, " ", "-")

	return output
}

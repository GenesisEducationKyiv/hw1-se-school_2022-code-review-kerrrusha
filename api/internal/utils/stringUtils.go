package utils

import "strings"

func RemoveRedundantGaps(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}

package utility

import "strings"

func escapeForLike(searchStr string) string {
	searchStr = strings.Replace(searchStr, "\\", "\\\\", -1)
	searchStr = strings.Replace(searchStr, "%", "\\%", -1)
	searchStr = strings.Replace(searchStr, "_", "\\_", -1)
	return searchStr
}

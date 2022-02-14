package util

import (
	"fmt"
	"strings"
)

// case-insensitive strings.contains
func StringContains(s string, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Contains(s, substr)
}

func FileHyperlink(file string) string {
	return Hyperlink("file://"+file, file)
}

func Hyperlink(url string, text string) string {
	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, text)
}

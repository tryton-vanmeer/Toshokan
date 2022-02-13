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

func FileHyperlink(file string, text string) string {
	return fmt.Sprintf("\033]8;;file://%s\033\\%s\033]8;;\033\\", file, text)
}
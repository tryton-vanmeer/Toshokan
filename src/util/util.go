package util

import "strings"

// case-insensitive strings.contains
func StringContains(s string, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Contains(s, substr)
}

package pkg

import (
	"strings"
)

func globalToLowerCase(s string) string {
	return strings.ToLower(s)
}

func globalTrimSpace(s string) string {
	return strings.TrimSpace(globalToLowerCase(s))
}

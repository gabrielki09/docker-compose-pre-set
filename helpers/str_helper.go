package helpers

import "strings"

func GlobalToLowerCase(s string) string { return strings.ToLower(s) }

func GlobalTrimSpace(s string) string { return strings.TrimSpace(GlobalToLowerCase(s)) }

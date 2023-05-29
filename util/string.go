package util

import "strings"

func StripQuotes(s string) string {
	return strings.Trim(s, `"'`)
}

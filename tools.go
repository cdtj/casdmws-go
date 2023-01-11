package casdmwsgo

import "strings"

// Concat concats strings in a fast way
func Concat(sep string, str ...string) string {
	return strings.Join(str, sep)
}

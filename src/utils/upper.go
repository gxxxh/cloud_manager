package utils

import "strings"

func UpperFirst(s string) string {
	first := strings.ToUpper(s[0:1])
	return first + s[1:]
}

func LowerFirst(s string) string {
	first := strings.ToLower(s[0:1])
	return first + s[1:]
}

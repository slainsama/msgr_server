package utils

import "strings"

// EscapeChar escapes special characters in a string
func EscapeChar(s string) string {
	// Escape '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=','|', '{', '}', '.', '!'  in a string with regex
	mathChar := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, char := range mathChar {
		s = strings.ReplaceAll(s, char, "\\"+char)
	}
	return s
}

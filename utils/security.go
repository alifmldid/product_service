package utils

import "html"

func SanitizeInput(input string) string {
	return html.EscapeString(input)
}

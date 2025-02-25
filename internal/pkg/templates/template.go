package templates

import (
	"strings"
)

func PrintCell(content string, maxLength int) string {
	contentLen := len(content)
	if contentLen%2 != 0 {
		content += " "
	}

	if maxLength%2 != 0 {
		maxLength += 1
	}

	paddingLen := (((maxLength - len(content)) / 2) + 2)
	padding := strings.Repeat(" ", paddingLen)
	return padding + content + padding + "|"
}

func PrintHR(maxLengths ...int) string {
	hr := "-"

	for _, maxLen := range maxLengths {
		if maxLen%2 != 0 {
			maxLen += 1
		}
		hr += strings.Repeat("-", maxLen+5)
	}

	return hr
}

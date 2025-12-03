package utils

import (
	"fmt"
	"strings"
)

// ReadSection extracts a section from a MarkDows-style text.
// section is a header string such as "## readme", "## report", or "## review".
// It returns the text following the given section header, excluding the header
// itself, up to the next "## " header or the end of the document.
func ReadSection(markdown string, header string) string {
	needle := strings.ToLower(fmt.Sprintf("## %s", header))
	start := strings.Index(strings.ToLower(markdown), needle)
	if start == -1 {
		return "" // not found
	}
	rest := markdown[start:]

	lines := strings.Split(rest, "\n")

	var out []string
	for _, l := range lines[1:] {
		if strings.HasPrefix(strings.ToLower(l), "## ") {
			break
		}
		out = append(out, l)
	}

	result := strings.Join(out, "\n")
	return result
}

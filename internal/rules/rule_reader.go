package rules

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed rules/RULES.md
var RulesMd string

// LoadRules reads a RULES.md file and extracts the specified section.
// If the "RULES_PATH" environment variable is set, it reads from
// "{RULES_PATH}/RULES.md", otherwise it falls back to "RULES.md".
func LoadRules(section string) (string, error) {
	var sectionStr string
	if path := os.Getenv("RULES_PATH"); path != "" {
		abs, err := filepath.Abs(path)
		if err != nil {
			sectionStr = readSection(RulesMd, section)
			return sectionStr, fmt.Errorf("your RULES_PATH is wrong %s, fallback to default rules", path)
		}

		file, err := os.ReadFile(abs)
		if err != nil {
			sectionStr = readSection(RulesMd, section)
			return sectionStr, errors.New("failed to load RULES_PATH file, fallback to default rules")
		}

		sectionStr = readSection(string(file), section)
		if sectionStr == "" {
			sectionStr = readSection(RulesMd, section)
			return sectionStr, errors.New("no rule section, fallback to default rules")
		}
		return sectionStr, nil
	}

	sectionStr = readSection(RulesMd, section)
	return sectionStr, nil
}

// readSection extracts a section from a RULES.md-style text.
// rules is the full content of RULES.md,
// section is a header string such as "## readme", "## report", or "## review".
// It returns the text following the given section header, excluding the header
// itself, up to the next "## " header or the end of the document.
func readSection(rules string, section string) string {
	needle := strings.ToLower(fmt.Sprintf("## %s", section))
	start := strings.Index(strings.ToLower(rules), needle)
	if start == -1 {
		return "" // not found
	}
	rest := rules[start:]

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

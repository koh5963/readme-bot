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

func readSection(rules string, section string) string {
	needle := strings.ToLower(fmt.Sprintf("## %s", section))
	start := strings.Index(strings.ToLower(rules), needle)
	if start == -1 {
		return "" // not found
	}
	// スタート位置から先だけ抽出
	rest := rules[start:]

	// rest を行 split
	lines := strings.Split(rest, "\n")

	var out []string
	for _, l := range lines[1:] { // 1行目は "## readme"
		if strings.HasPrefix(strings.ToLower(l), "## ") {
			break
		}
		out = append(out, l)
	}

	result := strings.Join(out, "\n")
	return result
}

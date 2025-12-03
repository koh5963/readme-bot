package rules

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/koh5963/readme-bot/internal/utils"
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
			sectionStr = utils.ReadSection(RulesMd, section)
			return sectionStr, fmt.Errorf("your RULES_PATH is wrong %s, fallback to default rules", path)
		}

		file, err := os.ReadFile(abs)
		if err != nil {
			sectionStr = utils.ReadSection(RulesMd, section)
			return sectionStr, errors.New("failed to load RULES_PATH file, fallback to default rules")
		}

		sectionStr = utils.ReadSection(string(file), section)
		if sectionStr == "" {
			sectionStr = utils.ReadSection(RulesMd, section)
			return sectionStr, errors.New("no rule section, fallback to default rules")
		}
		return sectionStr, nil
	}

	sectionStr = utils.ReadSection(RulesMd, section)
	return sectionStr, nil
}

package rules

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed rules/RULES.md
var RulesMd string

func LoadRules() (string, error) {
	if path := os.Getenv("RULES_PATH"); path != "" {
		abs, err := filepath.Abs(path)
		if err != nil {
			return RulesMd, fmt.Errorf("your RULES_PATH is wrong %s, fallback to embedded RULES.md", path)
		}

		file, err := os.ReadFile(abs)
		if err != nil {
			return RulesMd, errors.New("failed to load RULES_PATH file, fallback to embedded RULES.md")
		}
		return string(file), nil
	}
	return RulesMd, nil
}

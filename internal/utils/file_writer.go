package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RewriteReadme(content, header string) error {

	err := os.Rename("./README.md", "./_README.md")
	if err != nil {
		return fmt.Errorf("rename README file failed: %v", err)
	}

	oldFile, err := os.Open("./_README.md")
	if err != nil {
		_ = os.Rename("./_README.md", "./README.md")
		return fmt.Errorf("old new README file failed: %v", err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("./README.md")
	if err != nil {
		oldFile.Close()
		_ = os.Rename("./_README.md", "./README.md")
		return fmt.Errorf("create new README file failed: %v", err)
	}
	defer newFile.Close()

	isExistHeader := false
	scanner := bufio.NewScanner(oldFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == strings.TrimSpace(header) {
			isExistHeader = true
			fmt.Fprintln(newFile, header)
			fmt.Fprintln(newFile, content)
			continue
		}
		fmt.Fprintln(newFile, line)
	}
	if !isExistHeader {
		fmt.Fprintln(newFile, header)
		fmt.Fprintln(newFile, content)
	}

	os.Remove("./_README.md")
	return nil
}

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	oldStr := "github.com/jibankumarpanda/gogcli.1/"
	newStr := "github.com/jibankumarpanda/gogcli.1/"

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walk %q: %w", path, err)
		}

		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == ".tools" {
				return filepath.SkipDir
			}

			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".go" && d.Name() != "Makefile" {
			return nil
		}

		content, err := os.ReadFile(path) // #nosec G304 -- paths come from filepath.WalkDir over local repo tree
		if err != nil {
			return fmt.Errorf("read %q: %w", path, err)
		}

		if bytes.Contains(content, []byte(oldStr)) {
			newContent := bytes.ReplaceAll(content, []byte(oldStr), []byte(newStr))

			err = os.WriteFile(path, newContent, 0o600) // #nosec G306 -- utility script output
			if err != nil {
				return fmt.Errorf("write %q: %w", path, err)
			}

			println("Updated:", path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"bytes"
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
			return err
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

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if bytes.Contains(content, []byte(oldStr)) {
			newContent := bytes.ReplaceAll(content, []byte(oldStr), []byte(newStr))
			err = os.WriteFile(path, newContent, 0644)
			if err != nil {
				return err
			}
			println("Updated:", path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

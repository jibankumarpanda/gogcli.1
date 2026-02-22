package main

import (
	"os"

	"github.com/jibankumarpanda/gogcli.1/internal/cmd"
)

func main() {
	if err := cmd.Execute(os.Args[1:]); err != nil {
		os.Exit(cmd.ExitCode(err))
	}
}

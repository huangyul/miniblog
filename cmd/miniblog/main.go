package main

import (
	"miniblog/internal/miniblog"
	"os"
)

func main() {

	cmd := miniblog.NewMiniBlogCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

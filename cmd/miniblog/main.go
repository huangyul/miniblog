package main

import (
	_ "go.uber.org/automaxprocs"
	"miniblog/internal/miniblog"
	"os"
)

func main() {
	commoand := miniblog.NewMiniBlogCommand()
	if err := commoand.Execute(); err != nil {
		os.Exit(1)
	}
}

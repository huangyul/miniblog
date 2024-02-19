package main

import (
	"miniblog/internal/miniblog"
)

func main() {
	cmd := miniblog.NewMiniBlogCommand()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

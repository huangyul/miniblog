package main

import "miniblog/internal/miniblog"

func main() {
	cmd := miniblog.NewMiniblogCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

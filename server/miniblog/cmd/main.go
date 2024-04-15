package main

import "github.com/huangyul/miniblog/internal/miniblog"

func main() {
	cmd := miniblog.NewMiniBlogCommand()
	cmd.Execute()
}

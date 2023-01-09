package main

import (
	"go-zip/pkg/runner"
)

func main() {
	options := runner.ParseOptions()
	r := runner.NewRunner(options)
	r.Run()
}

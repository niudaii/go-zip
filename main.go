package main

import (
	"context"
	"go-zip/pkg/runner"
)

func main() {
	options := runner.ParseOptions()
	newRunner := runner.NewRunner(options)
	newRunner.Run(context.Background())
}

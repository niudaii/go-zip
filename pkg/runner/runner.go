package runner

import (
	"fmt"
	"go-zip/pkg/pack"
	"os"
	"time"
)

type Runner struct {
	options *Options
}

func NewRunner(options *Options) *Runner {
	runner := &Runner{
		options: options,
	}
	return runner
}

func (r *Runner) Run() {
	start := time.Now()
	if r.options.List {
		pack.ListFileType(r.options.Dir)
		os.Exit(0)
	}
	if err := pack.Zip(r.options.Dir, r.options.Output, r.options.BlackExt, r.options.WhiteExt); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("保存位置: %v\n", r.options.Output)
	fmt.Printf("运行时间: %v\n", time.Since(start))
}

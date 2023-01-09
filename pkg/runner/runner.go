package runner

import (
	"go-zip/pkg/pack"
	"log"
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
	} else {
		err := pack.Zip(r.options.Dir, r.options.Output, r.options.BlackExt, r.options.WhiteExt)
		if err != nil {
			log.Println(err)
		}
		log.Printf("保存位置: %v\n", r.options.Output)
	}
	log.Printf("运行时间: %v\n", time.Since(start))
}

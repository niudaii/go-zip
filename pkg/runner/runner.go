package runner

import (
	"context"
	"fmt"
	"go-zip/pkg/oss"
	"go-zip/pkg/pack"
	"go-zip/pkg/util"
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

func (r *Runner) Run(ctx context.Context) {
	start := time.Now()
	if r.options.List {
		pack.ListFileType(r.options.SrcFile)
		os.Exit(0)
	}
	err := pack.Zip(r.options.SrcFile, r.options.DestZip, r.options.Exclude)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("[INF] Saved: %v\n", r.options.DestZip)
	if r.options.Upload {
		objName := util.GetRandomString(6)
		oss.UploadFile(r.options.DestZip, objName)
		fmt.Printf("[INF] OSS-Obj: %v\n", objName)
	}
	fmt.Printf("[INF] Spent: %v", time.Since(start))
}

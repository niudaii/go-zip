package runner

import (
	"flag"
	"fmt"
)

type Options struct {
	Dir     string
	Output  string
	Exclude string
	List    bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.Dir, "dir", "./", "dir to pack")
	flag.StringVar(&options.Output, "output", "output.zip", "file to write output")
	flag.BoolVar(&options.List, "list", false, "list file type")
	flag.StringVar(&options.Exclude, "exclude", ".css,.js,.jpg,.jpeg,.png,.gif,.bmp,.tiff,.tif,.psd,.xcf", "exclude suffix name")
	flag.Parse()
	showBanner()
	options.validateOptions()
	return options
}

func (options *Options) validateOptions() {
	fmt.Printf("打包路径: %v\n", options.Dir)
}

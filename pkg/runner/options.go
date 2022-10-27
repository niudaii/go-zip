package runner

import (
	"flag"
	"fmt"
)

type Options struct {
	Dir      string
	Output   string
	BlackExt string
	WhiteExt string
	List     bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.Dir, "d", "./", "dir to pack")
	flag.StringVar(&options.Output, "o", "output.zip", "file to write output")
	flag.BoolVar(&options.List, "l", false, "list file type")
	flag.StringVar(&options.BlackExt, "black", ".css,.js,.jpg,.jpeg,.png,.gif,.bmp", "black suffix names[.css]")
	flag.StringVar(&options.WhiteExt, "white", "", "white suffix name, [.php,.go]")
	flag.Parse()
	showBanner()
	options.validateOptions()
	return options
}

func (options *Options) validateOptions() {
	fmt.Printf("打包路径: %v\n", options.Dir)
}

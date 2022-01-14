package runner

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	SrcFile string
	DestZip string
	Exclude string
	Upload  bool
	List    bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.SrcFile, "src", "./", "src file")
	flag.StringVar(&options.DestZip, "dest", "", "dest zip")
	flag.BoolVar(&options.List, "list", false, "list file type")
	flag.StringVar(&options.Exclude, "exclude", ".css,.js", "exclude suffix name")
	flag.BoolVar(&options.Upload, "upload", false, "upload the file to oss")
	flag.Parse()
	showBanner()
	options.validateOptions()
	return options
}

func (options *Options) validateOptions() {
	if !options.List {
		if options.DestZip == "" {
			fmt.Println("please input dest zip")
			os.Exit(1)
		}
	}
}

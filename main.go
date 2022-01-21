package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

//
var (
    arg_help bool
	arg_version bool
    arg_conf string
	arg_wd string
)

func Init() {
    flag.BoolVar(&arg_help, "h", false, "Help with the usage")
    flag.BoolVar(&arg_version, "v", false, "Show version and exit")
    flag.StringVar(&arg_conf, "c", "~/.go-perspective.conf", "set `config` file")
	flag.StringVar(&arg_wd, "d", ".", "Set work directory")

    // Usage
    flag.Usage = usage
}

func main() {
	Init()

	flag.Parse()

	if arg_help {
		flag.Usage();
		return;
	}

	if arg_version {
		version();
		return;
	}

	//fmt.Println(arg_wd)
	getFilelist(arg_wd)
}

func usage() {
    fmt.Fprintf(os.Stderr, `Usage: go-perspective`)
	fmt.Fprintf(os.Stderr, "\n\nOptions:\n")
    flag.PrintDefaults()
}

func version() {
	fmt.Fprintf(os.Stderr, "Version: go-perspective: 0.1.0\n")
}

func getFilelist(tpath string) {
	err := filepath.Walk(tpath, func(path string, f os.FileInfo, err error) error {

		if ( f == nil ) {return err}

		if f.IsDir() {

		}else{
			fmt.Println(path)
			// fmt.Println(f.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

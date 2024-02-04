package main

import (
	"fmt"
	"os"

	"github.com/Frank-Mayer/list/internal/list"
	"github.com/Frank-Mayer/list/internal/version"
	"github.com/alecthomas/kingpin/v2"
)

var (
	paths = kingpin.Arg("paths", "the files(s) and/or folder(s) to display").Default(".").Strings()
	all   = kingpin.Flag("all", "show hidden files").Short('a').Bool()
	table = kingpin.Flag("table", "list in table format").Short('l').Bool()
	tree  = kingpin.Flag("tree", "list in tree format").Short('t').Bool()
	count = kingpin.Flag("count", "count the files").Short('c').Bool()
	ves   = kingpin.Flag("version", "show version").Short('v').Bool()
)

func main() {
	kingpin.Parse()

	if *ves {
		fmt.Println(version.Version)
		return
	}

	fn := list.Default

	if *table {
		fn = list.Table
	}
	if *tree {
		fn = list.Tree
	}
	if *count {
		fn = list.Count
	}

	options := &list.Options{
		Hidden: *all,
	}

	var err error
	switch len(*paths) {
	case 0:
		err = fn(".", options)
	case 1:
		err = fn((*paths)[0], options)
	default:
		for i, arg := range *paths {
			if i > 0 {
				println()
			}
			fmt.Println(arg + ":")
			err = fn(arg, options)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

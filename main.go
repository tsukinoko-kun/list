package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Frank-Mayer/list/internal/list"
	"github.com/Frank-Mayer/list/internal/version"
)

var (
	table = flag.Bool("l", false, "list in table format")
	all   = flag.Bool("a", false, "list hidden files")
	tree  = flag.Bool("t", false, "list in tree format")
	ves   = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

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

	options := &list.Options{
		Hidden: *all,
	}

	var err error
	restArgs := flag.Args()
	switch len(restArgs) {
	case 0:
		err = fn(".", options)
	case 1:
		err = fn(restArgs[0], options)
	default:
		for i, arg := range restArgs {
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

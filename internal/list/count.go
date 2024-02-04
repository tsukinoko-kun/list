package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func Count(p string, options *Options) error {
	// check if path is a directory
	fi, err := os.Stat(p)
	if err != nil {
		return errors.Wrap(err, "could not get file info for path "+p)
	}

	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}

	// get the list of files
	files, err := os.ReadDir(p)
	if err != nil {
		return errors.Wrap(err, "could not read directory "+p)
	}

	// count the files
	count := 0
	for _, file := range files {
		if !options.Hidden && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		count++
	}

	fmt.Println(count)
	return nil
}

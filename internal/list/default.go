package list

import (
	"os"
	"path/filepath"

	"github.com/Frank-Mayer/list/internal/utils"
	"github.com/pkg/errors"
)

func Default(p string, options *Options) error {
	// check if path is a directory
	fi, err := os.Stat(p)
	if err != nil {
		return errors.Wrap(err, "could not get file info for path "+p)
	}

	if !fi.IsDir() {
		err = printStyled(p)
		if err != nil {
			return err
		}
		println()
		return nil
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		absEntryPath := filepath.Join(p, entry.Name())
		hidden, err := utils.IsHiddenFile(absEntryPath)
		if err != nil {
			return errors.Wrap(err,"could not check if file is hidden at path "+absEntryPath)
		}
		if !options.Hidden && hidden {
			continue
		}
		err = printStyled(absEntryPath)
		if err != nil {
			return err
		}
		print(segment)
	}

	println()

	return nil
}

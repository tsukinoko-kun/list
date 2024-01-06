package list

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Frank-Mayer/list/internal/utils"
)

func Default(p string, options *Options) error {
	// check if path is a directory
	fi, err := os.Stat(p)
	if err != nil {
		return errors.Join(errors.New("could not get file info for path "+p), err)
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
			return errors.Join(errors.New("could not check if file is hidden at path "+absEntryPath), err)
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

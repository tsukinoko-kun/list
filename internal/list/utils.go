package list

import (
	"os"
	"path/filepath"

	"github.com/tsukinoko-kun/list/internal/utils"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

const (
	segment    = "  "
	executable = 0b001001001
)

func printStyled(p string) error {
	content, c, err := printStyle(p)
	if err != nil {
		return err
	}
	c.Print(content)
	return nil
}

func printStyle(p string) (string, *color.Color, error) {
	content := filepath.Base(p)
	c := cFile

	isExec := false

	fi, err := os.Lstat(p)
	if err != nil {
		return content, c, errors.Wrap(err,"could not get file info for path "+p)
	}

	if fi.IsDir() {
		c = cDir
		content += "/"
	} else {
		perm := fi.Mode().Perm()
		if perm&executable != 0 {
			isExec = true
		}
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		c = cLink
		content += "@"
	}

	hidden, err := utils.IsHiddenFile(p)
	if err != nil {
		return content, c, errors.Wrap(err, "could not check if file is hidden for path "+p)
	}
	if hidden {
		c = cHidden
	}

	if isExec {
		content = cExec.Sprint("*") + content
	}

	return content, c, nil
}

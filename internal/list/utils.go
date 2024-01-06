package list

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Frank-Mayer/list/internal/utils"
	"github.com/fatih/color"
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
		return content, c, errors.Join(errors.New("could not get file info for path "+p), err)
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
		return content, c, errors.Join(errors.New("could not check if file is hidden for path "+p), err)
	}
	if hidden {
		c = cHidden
	}

	if isExec {
		cExec.Print("*")
	}

	return content, c, nil
}

package list

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Frank-Mayer/list/internal/utils"
)

const (
	segment    = "  "
	executable = 0b001001001
)

func printColored(p string) error {
	content := filepath.Base(p)
	color := cFile

	isExec := false

	fi, err := os.Lstat(p)
	if err != nil {
		return errors.Join(errors.New("could not get file info for path "+p), err)
	}

	if fi.IsDir() {
		color = cDir
		content += "/"
	} else {
		perm := fi.Mode().Perm()
		if perm&executable != 0 {
			isExec = true
		}
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		color = cLink
		content += "@"
	}

	hidden, err := utils.IsHiddenFile(p)
	if err != nil {
		return errors.Join(errors.New("could not check if file is hidden for path "+p), err)
	}
	if hidden {
		color = cHidden
	}

	color.Print(content)
	if isExec {
		cExec.Print("*")
	}

	return nil
}

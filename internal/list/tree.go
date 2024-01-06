package list

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Frank-Mayer/list/internal/utils"
)

const (
	osPathSep = string(os.PathSeparator)
)

func Tree(p string, options *Options) error {
	abs, err := filepath.Abs(p)
	if err != nil {
		return errors.Join(errors.New("could not get absolute path for tree of path "+p), err)
	}
	rootLength := len(strings.Split(abs, osPathSep))

	err = walk(p, options.Hidden, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Join(errors.New("to level walk function passed error"), err)
		}

		hidden, err := utils.IsHiddenFile(path)
		if err != nil {
			return errors.Join(errors.New("could not check if file is hidden for tree at path "+path), err)
		}
		if !options.Hidden && hidden {
			return nil
		}

		length := len(strings.Split(path, osPathSep)) - rootLength
		print(repeat(segment, length))

		err = printStyled(path)
		if err != nil {
			return errors.Join(errors.New("could not print colored path for tree at path "+path), err)
		}

		println()
		return nil
	})

	if err != nil {
		return errors.Join(errors.New("could not walk path for tree at path "+p), err)
	}

	return nil
}

func repeat(s string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

func walk(root string, all bool, fn filepath.WalkFunc) error {
	q := utils.NewStack[string]()
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return errors.Join(errors.New("could not get absolute path for walk of path "+root), err)
	}
	q.Push(absRoot)

	for qEntry := q.Pop(); qEntry != nil; qEntry = q.Pop() {
		fi, err := os.Lstat(*qEntry)
		if err != nil {
			return errors.Join(errors.New("could not get file info in walk for path "+*qEntry), err)
		}

		if !all {
			hidden, err := utils.IsHiddenFile(*qEntry)
			if err != nil {
				return errors.Join(errors.New("could not check if file is hidden for walk at path "+*qEntry), err)
			}
			if hidden {
				continue
			}
		}

		if fi.IsDir() {
			entries, err := os.ReadDir(*qEntry)
			if err != nil {
				// if permission denied, continue
				if os.IsPermission(err) {
					continue
				}
				return errors.Join(errors.New("could not read directory "+*qEntry), err)
			}

			var dirEntry fs.DirEntry
			for i := len(entries) - 1; i >= 0; i-- {
				dirEntry = entries[i]
				path := filepath.Join(*qEntry, dirEntry.Name())
				q.Push(path)
			}
		}

		err = fn(*qEntry, fi, nil)
		if err != nil {
			return errors.Join(errors.New("could not execute walk file tree function for path "+*qEntry), err)
		}
	}

	return nil
}

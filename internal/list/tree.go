package list

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/tsukinoko-kun/list/internal/utils"
	"github.com/pkg/errors"
)

const (
	osPathSep = string(os.PathSeparator)
)

func Tree(p string, options *Options) error {
	abs, err := filepath.Abs(p)
	if err != nil {
		return errors.Wrap(err, "could not get absolute path for tree of path "+p)
	}
	rootLength := len(strings.Split(abs, osPathSep))

	err = walk(p, options.Hidden, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "to level walk function passed error")
		}

		hidden, err := utils.IsHiddenFile(path)
		if err != nil {
			return errors.Wrap(err, "could not check if file is hidden for tree at path "+path)
		}
		if !options.Hidden && hidden {
			return nil
		}

		length := len(strings.Split(path, osPathSep)) - rootLength
		print(repeat(segment, length))

		err = printStyled(path)
		if err != nil {
			return errors.Wrap(err, "could not print colored path for tree at path "+path)
		}

		println()
		return nil
	})

	if err != nil {
		return errors.Wrap(err, "could not walk path for tree at path "+p)
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
		return errors.Wrap(err, "could not get absolute path for walk of path "+root)
	}
	q.Push(absRoot)

	for qEntry := q.Pop(); qEntry != nil; qEntry = q.Pop() {
		fi, err := os.Lstat(*qEntry)
		if err != nil {
			return errors.Wrap(err, "could not get file info in walk for path "+*qEntry)
		}

		if !all {
			hidden, err := utils.IsHiddenFile(*qEntry)
			if err != nil {
				return errors.Wrap(err, "could not check if file is hidden for walk at path "+*qEntry)
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
				return errors.Wrap(err, "could not read directory "+*qEntry)
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
			return errors.Wrap(err, "could not execute walk file tree function for path "+*qEntry)
		}
	}

	return nil
}

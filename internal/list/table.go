package list

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/tsukinoko-kun/list/internal/utils"
	"github.com/pkg/errors"
)

func Table(p string, options *Options) error {
	fi, err := os.Stat(p)
	if err != nil {
		return err
	}

	table := utils.NewTable(7)

	if fi.IsDir() {
		if options.Hidden {
			te, err := tableEntry(table, true, false, fi)
			if err != nil {
				return errors.Wrap(err,"could not create table entry for path "+p)
			}
			err = te.AddCell("./")
			if err != nil {
				return errors.Wrap(err, "could not add cell to table at path "+p)
			}
			te.SetStyle(cDir)
		}
	} else {
		te, err := tableEntry(table, false, false, fi)
		if err != nil {
			return errors.Wrap(err, "could not create table entry for path "+p)
		}
		err = te.AddCell(filepath.Base(p))
		if err != nil {
			return errors.Wrap(err, "could not add cell to table at path "+p)
		}
		te.SetStyle(cFile)
		return nil
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		return errors.Wrap(err, "could not read directory "+p)
	}

	for _, entry := range entries {
		hidden, err := utils.IsHiddenFile(entry.Name())
		if err != nil {
			return errors.Wrap(err, "could not check if file is hidden")
		}
		if !options.Hidden && hidden {
			continue
		}

		name := entry.Name()
		isDir := entry.IsDir()
		isSymlink := entry.Type()&fs.ModeSymlink != 0
		absEntryPath := filepath.Join(p, name)
		fi, err := entry.Info()
		if err != nil {
			return errors.Wrap(err, "could not get file info for path "+absEntryPath)
		}

		te, err := tableEntry(table, isDir, isSymlink, fi)
		if err != nil {
			return errors.Wrap(err, "could not create table entry for path "+absEntryPath)
		}

		desplayName, style, err := printStyle(absEntryPath)
		if err != nil {
			return errors.Wrap(err, "could not print colored path for table at path "+absEntryPath)
		}
		err = te.AddCell(desplayName)
		if err != nil {
			return errors.Wrap(err, "could not add cell to table at path "+absEntryPath)
		}
		te.SetStyle(style)
	}

	table.Print()

	return nil
}

func permsString(isDir bool, isSymlink bool, fi fs.FileInfo) string {
	str := make([]byte, 13)
	if isDir {
		str[0] = 'd'
	} else if isSymlink {
		str[0] = 'l'
	} else {
		str[0] = '-'
	}

	perms := fi.Mode().Perm()

	for i := 2; i >= 0; i-- {
		localPerms := perms >> (i * 3) & 0b111
		if localPerms&0b100 != 0 {
			str[1+(3-i)*3] = 'r'
		} else {
			str[1+(3-i)*3] = '-'
		}
		if localPerms&0b010 != 0 {
			str[2+(3-i)*3] = 'w'
		} else {
			str[2+(3-i)*3] = '-'
		}
		if localPerms&0b001 != 0 {
			str[3+(3-i)*3] = 'x'
		} else {
			str[3+(3-i)*3] = '-'
		}
	}

	return string(str)
}

func tableEntry(t *utils.Table, isDir bool, isSymlink bool, fi fs.FileInfo) (*utils.TableEntry, error) {
	te := t.NewEntry()
	var err error
	err = te.AddCell(permsString(isDir, isSymlink, fi))
	if err != nil {
		return nil, errors.Wrap(err, "could not add cell to table")
	}

	username, groupname, nlink, err := utils.Owner(fi)
	if err == nil {
		err = te.AddCell(username)
		if err != nil {
			return nil, errors.Wrap(err, "could not add cell to table")
		}
		err = te.AddCell(groupname)
		if err != nil {
			return nil, errors.Wrap(err, "could not add cell to table")
		}
		err = te.AddCell(nlink)
		if err != nil {
			return nil, errors.Wrap(err, "could not add cell to table")
		}
	}
	err = te.AddCell(utils.FormatBytes(fi.Size()))
	if err != nil {
		return nil, errors.Wrap(err, "could not add cell to table")
	}
	err = te.AddCell(fi.ModTime().Format(time.DateTime))
	if err != nil {
		return nil, errors.Wrap(err, "could not add cell to table")
	}

	return te, nil
}

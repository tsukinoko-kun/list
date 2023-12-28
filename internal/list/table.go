package list

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Frank-Mayer/list/internal/utils"
)

func Table(p string, options *Options) error {
	fi, err := os.Stat(p)
	if err != nil {
		return err
	}

	fmt.Printf("total %s\n", utils.FormatBytes(fi.Size()))

	if fi.IsDir() {
		if options.Hidden {
			currentDirString, err := tableEntry(p, true, false, fi)
			if err != nil {
				return err
			}
			print(currentDirString)
			cDir.Println("./")
		}
	} else {
		currentDirString, err := tableEntry(p, false, false, fi)
		if err != nil {
			return err
		}
		print(currentDirString)
		cFile.Println(filepath.Base(p))
		return nil
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		return errors.Join(errors.New("could not read directory "+p), err)
	}
	for _, entry := range entries {
		hidden, err := utils.IsHiddenFile(entry.Name())
		if err != nil {
			return errors.Join(errors.New("could not check if file is hidden"), err)
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
			return errors.Join(errors.New("could not get file info for path "+absEntryPath), err)
		}

		tableEntryString, err := tableEntry(absEntryPath, isDir, isSymlink, fi)
		if err != nil {
			return errors.Join(errors.New("could not get table entry for path "+absEntryPath), err)
		}

		print(tableEntryString)
		err = printColored(absEntryPath)
		if err != nil {
			println(name)
			return errors.Join(errors.New("could not print colored path for table at path "+absEntryPath), err)
		}
		println()
	}

	return nil
}

func tableEntry(p string, isDir bool, isSymlink bool, fi fs.FileInfo) (string, error) {
	permsString := make([]byte, 13)
	if isDir {
		permsString[0] = 'd'
	} else if isSymlink {
		permsString[0] = 'l'
	} else {
		permsString[0] = '-'
	}

	perms := fi.Mode().Perm()

	for i := 2; i >= 0; i-- {
		localPerms := perms >> (i * 3) & 0b111
		if localPerms&0b100 != 0 {
			permsString[1+(3-i)*3] = 'r'
		} else {
			permsString[1+(3-i)*3] = '-'
		}
		if localPerms&0b010 != 0 {
			permsString[2+(3-i)*3] = 'w'
		} else {
			permsString[2+(3-i)*3] = '-'
		}
		if localPerms&0b001 != 0 {
			permsString[3+(3-i)*3] = 'x'
		} else {
			permsString[3+(3-i)*3] = '-'
		}
	}

	stat, ok := fi.Sys().(*syscall.Stat_t)
	sb := strings.Builder{}
	sb.Write(permsString)
	if ok {
		username := fmt.Sprintf("%d", stat.Uid)
		usr, err := user.LookupId(username)
		if err == nil {
			username = usr.Username
		}

		groupname := fmt.Sprintf("%d", stat.Gid)
		group, err := user.LookupGroupId(groupname)
		if err == nil {
			groupname = group.Name
		}

		sb.WriteString(fmt.Sprintf(" %4d %s %s", stat.Nlink, username, groupname))
	}
	sb.WriteString(fmt.Sprintf(" %12s ", utils.FormatBytes(fi.Size())))
	sb.WriteString(fi.ModTime().Format(time.RFC3339))
	sb.WriteRune(' ')

	return sb.String(), nil
}

//go:build !windows

package utils

import (
	"fmt"
	"io/fs"
	"os/user"
	"syscall"
)

func Owner(fi fs.FileInfo) (string, string, string, error) {
	var username, groupname, nlink string

	stat, ok := fi.Sys().(*syscall.Stat_t)
	if ok {
		username = fmt.Sprintf("%d", stat.Uid)
		usr, err := user.LookupId(username)
		if err == nil {
			username = usr.Username
		}

		groupname = fmt.Sprintf("%d", stat.Gid)
		group, err := user.LookupGroupId(groupname)
		if err == nil {
			groupname = group.Name
		}

		nlink = fmt.Sprintf("%d", stat.Nlink)
	}

	return nlink, username, groupname, nil
}

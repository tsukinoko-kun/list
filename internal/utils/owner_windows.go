//go:build windows

package utils

import (
	"io/fs"
)

func Owner(fi fs.FileInfo) string {
	return ""
}

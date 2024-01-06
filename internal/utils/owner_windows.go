//go:build windows

package utils

import (
	"errors"
	"io/fs"
)

func Owner(fi fs.FileInfo) (string, string, string, error) {
	return "", "", "", errors.New("not implemented")
}

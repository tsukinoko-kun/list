//go:build !windows

package utils

import "path/filepath"

func IsHiddenFile(p string) (bool, error) {
	filename := filepath.Base(p)
	return filename[0] == '.', nil
}

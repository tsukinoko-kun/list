//go:build windows

package utils

import (
	"syscall"
)

func IsHiddenFile(p string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(p)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}

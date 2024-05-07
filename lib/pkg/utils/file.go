package utils

import "os"

func IsFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

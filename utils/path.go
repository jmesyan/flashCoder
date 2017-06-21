package utils

import (
	"path/filepath"
	"strings"
)

func GetRootDirectory() string {
	dir, err := filepath.Abs(filepath.Dir("./"))
	CheckError(err)
	return strings.Replace(dir, "\\", "/", -1)
}

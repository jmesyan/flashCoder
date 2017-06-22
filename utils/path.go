package utils

import (
	"path/filepath"
	"strings"
)

func GetRootDirectory() string {
	dir, err := filepath.Abs(filepath.Dir("./"))
	LogError("error", err)
	return strings.Replace(dir, "\\", "/", -1)
}

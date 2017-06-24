package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetRootDirectory() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("the utils path is not ok")
		os.Exit(1)
	}
	utilPath := filepath.Dir(file)
	rootPath := filepath.Dir(utilPath)
	return strings.Replace(rootPath, "\\", "/", -1)
}

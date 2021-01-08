package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}

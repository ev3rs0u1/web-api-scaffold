package dirutil

import (
	"os"
	"path/filepath"
)

func DirCreate(dir string) error {
	if DirExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

// DirExist checks if directory exist
func DirExist(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func JoinCurrentDir(p string) string {
	dir, err := os.Executable()
	if err != nil {
		return os.TempDir()
	}

	return filepath.Join(filepath.Dir(dir), p)
}

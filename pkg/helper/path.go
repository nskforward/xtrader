package helper

import (
	"os"
	"path"
	"path/filepath"
)

func ProcessDir() string {
	dir, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path.Dir(dir)
}

func ResolvePath(names ...string) string {
	names = append([]string{ProcessDir()}, names...)
	return filepath.Join(names...)
}

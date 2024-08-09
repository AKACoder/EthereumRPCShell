package utils

import (
	"os"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

func SaveFile(path string, content []byte) {
	_ = os.WriteFile(path, content, 0777)
}

func CreateDir(path string) {
	_ = os.MkdirAll(path, 0777)
}

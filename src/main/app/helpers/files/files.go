package files

import "os"

func Exist(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

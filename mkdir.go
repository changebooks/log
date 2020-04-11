package log

import "os"

func Mkdir(path string) error {
	if IsDirExist(path) {
		return nil
	}

	if err := os.MkdirAll(path, FilePerm); err != nil {
		return err
	}

	return nil
}

func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

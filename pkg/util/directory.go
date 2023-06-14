package util

import (
	"os"
)

func TryToMkdir(path string, mode os.FileMode) error {
	var err error
	if ok, _ := PathExists(path); !ok {
		err = os.MkdirAll(path, mode)
	}

	return err
}

package go_utils

import "os"

func MkDirByPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return err
}

package go_utils

import "os"

func BaseByFileName(fileName string) string {
	for i := len(fileName) - 1; i >= 0 && fileName[i] != '/'; i-- {
		if fileName[i] == '.' {
			return fileName[:i]
		}
	}
	return ""
}

func AppendContent(filePath string, msg string) error {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()
	_, err = logFile.WriteString(msg + " \n")
	if err != nil {
		return err
	}
	return nil
}

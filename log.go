package go_utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logging struct {
	sync.Mutex
	RootPath string // 保存路径 如 log
	Name     string // 日志项目名 如 mysql
	MaxSize  int64  // 单个日志文件的最大大小（MB）
}

func NewLog(RootPath, Name string, MaxSize int64) Logging {
	return Logging{
		RootPath: RootPath,
		Name:     Name,
		MaxSize:  MaxSize,
	}
}

func (l *Logging) writelog(pre string, msg string) error {
	l.Lock()
	defer l.Unlock()
	getwd, err := os.Getwd()
	if err != nil {
		return err
	}
	now := time.Now()
	dateStr := Date("Y-m-d", now)
	timeStr := Date("H:i:s", now)
	msg = timeStr + " " + msg
	dir := getwd + "/" + l.RootPath + "/" + pre
	err = MkDirByPath(dir)
	if err != nil {
		return err
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	today_max := 0
	for _, file := range files {
		if file.IsDir() {
		} else {
			base := BaseByFileName(file.Name())
			date_n := strings.Split(base, "_")
			if date_n[0] == dateStr {
				atoi, _ := strconv.Atoi(date_n[1])
				if atoi > today_max {
					today_max = atoi
				}
			}
		}
	}
	if today_max == 0 {
		filePath := dir + "/" + dateStr + "_1.log"
		err := AppendContent(filePath, msg)
		if err != nil {
			return err
		}
	} else {
		oldPath := dir + "/" + dateStr + "_" + strconv.Itoa(today_max) + ".log"
		fileInfo, err := os.Stat(oldPath)
		if err != nil {
			return err
		}
		if fileInfo.Size() > l.MaxSize*1024*1024 {
			newPath := dir + "/" + dateStr + "_" + strconv.Itoa(today_max+1) + ".log"
			err := AppendContent(newPath, msg)
			if err != nil {
				return err
			}
		} else {
			err := AppendContent(oldPath, msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *Logging) Info(msg any) error {
	err := l.writelog(l.Name+"/info", fmt.Sprint(msg))
	if err != nil {
		return err
	}
	return nil
}

func (l *Logging) Warn(msg any) error {
	err := l.writelog(l.Name+"/warn", fmt.Sprint(msg))
	if err != nil {
		return err
	}
	return nil
}

func (l *Logging) Error(msg any) error {
	err := l.writelog(l.Name+"/error", fmt.Sprint(msg))
	if err != nil {
		return err
	}
	return nil
}

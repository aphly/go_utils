package go_utils

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logging struct {
	sync.Mutex
	RootPath string // 保存路径
	Name     string // 日志项目名
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
		AppendContent(filePath, msg)
	} else {
		oldPath := dir + "/" + dateStr + "_" + strconv.Itoa(today_max) + ".log"
		fileInfo, err := os.Stat(oldPath)
		if err != nil {
			return err
		}
		if fileInfo.Size() > l.MaxSize*1024*1024 {
			newPath := dir + "/" + dateStr + "_" + strconv.Itoa(today_max+1) + ".log"
			AppendContent(newPath, msg)
		} else {
			AppendContent(oldPath, msg)
		}
	}
	return nil
}

func (l *Logging) Info(args ...string) {
	if len(args) == 1 {
		l.writelog(l.Name+"/info", args[0])
	} else if len(args) == 2 {
		l.writelog(args[0]+"/info", args[1])
	}
}

func (l *Logging) Warn(args ...string) {
	if len(args) == 1 {
		l.writelog(l.Name+"/warn", args[0])
	} else if len(args) == 2 {
		l.writelog(args[0]+"/warn", args[1])
	}
}

func (l *Logging) Error(args ...string) {
	if len(args) == 1 {
		l.writelog(l.Name+"/error", args[0])
	} else if len(args) == 2 {
		l.writelog(args[0]+"/error", args[1])
	}
}

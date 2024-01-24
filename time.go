package go_utils

import (
	"math"
	"strings"
	"time"
)

func Date(format string, t time.Time) string {
	format = strings.Replace(format, "Y", "2006", 1)
	format = strings.Replace(format, "m", "01", 1)
	format = strings.Replace(format, "d", "02", 1)
	format = strings.Replace(format, "H", "15", 1)
	format = strings.Replace(format, "i", "04", 1)
	format = strings.Replace(format, "s", "05", 1)
	return t.Format(format)
}

func Int64toDate(sec int64) string {
	t := time.Unix(sec, 0)
	return Date("Y-m-d H:i:s", t)
}

func StringToTime(format string, str string) (time.Time, error) {
	format = strings.Replace(format, "Y", "2006", 1)
	format = strings.Replace(format, "m", "01", 1)
	format = strings.Replace(format, "d", "02", 1)
	format = strings.Replace(format, "H", "15", 1)
	format = strings.Replace(format, "i", "04", 1)
	format = strings.Replace(format, "s", "05", 1)
	tt, err := time.Parse(format, str)
	return tt, err
}

func CalculateDays(startDate string, endDate string) int {
	layout := "2006-01-02"
	tStart, err := time.Parse(layout, startDate)
	if err != nil {
		return -1
	}
	tEnd, err := time.Parse(layout, endDate)
	if err != nil {
		return -1
	}
	diff := tEnd.Sub(tStart).Hours() / 24
	return int(math.Abs(float64(diff)))
}

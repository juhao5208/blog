package logfile

import (
	"bytes"
	"fmt"
	"github.com/petermattis/goid"
	"github.com/sirupsen/logrus"
	"strconv"
)

// LogFormater 日志格式结构体
type LogFormater struct {
}

// Format 实现自定义的日志格式打印
func (l *LogFormater) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	newLog := fmt.Sprintf("%s %s msgid:%v %s\n", timestamp, entry.Level, l.getGoId, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

// getGoId 获取协程id
func (l *LogFormater) getGoId() string {
	id := strconv.Itoa(int(goid.Get()))
	if id == "" {
		return "unknow"
	}
	return id
}

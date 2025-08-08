package logfile

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"blog/consts"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*logrus.Logger
}

var logInstance *logrus.Logger // 日志实例
var logOnce sync.Once

// Instance 实例化
func Instance() *logrus.Logger {
	logOnce.Do(func() {
		if logInstance == nil {
			logInstance = logrus.New()
		}
	})
	return logInstance
}

// InitLogfile 初始化
func InitLogfile() {
	logInstance = Instance()
	logInstance.SetReportCaller(true)

	if consts.LogToStdout {
		logInstance.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logInstance.SetFormatter(&LogFormater{})
	}

	if consts.LogToStdout {
		logInstance.SetOutput(os.Stdout)
	} else {
		l := &lumberjack.Logger{
			Filename:   path.Join(consts.LogFilePath, consts.LogFileName),
			MaxSize:    consts.LogMaxSize,
			MaxBackups: consts.LogMaxBackup,
			Compress:   consts.LogCompress,
		}
		logInstance.SetOutput(l)
	}

	// 设置日志级别
	switch consts.LogLevel {
	case "debug":
		logInstance.SetLevel(logrus.DebugLevel)
	case "info":
		logInstance.SetLevel(logrus.InfoLevel)
	case "warn":
		logInstance.SetLevel(logrus.WarnLevel)
	case "error":
		logInstance.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logInstance.SetLevel(logrus.FatalLevel)
	case "panic":
		logInstance.SetLevel(logrus.PanicLevel)
	default:
		fmt.Printf("not support log level: %s\n", consts.LogLevel)
		os.Exit(1)
	}
}

func Trace(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(args...)
		}
	}()
	Instance().Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	Instance().Tracef(format, args...)
}

func Debug(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(args...)
		}
	}()
	Instance().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	Instance().Debugf(format, args...)
}

func Info(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(args...)
		}
	}()
	Instance().Info(args...)
}

func Infof(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	Instance().Infof(format, args...)
}

func Warn(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(args...)
		}
	}()
	Instance().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	Instance().Warnf(format, args...)
}

func Error(args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(args...)
		}
	}()

	tFilePath := ""
	for i := 5; i > 0; i-- {
		_, tfp, tno, ok := runtime.Caller(i)
		if !ok {
			continue
		}

		tFilePath += fmt.Sprintf("%s:%d -> ", path.Base(tfp), tno)
	}
	logstr := tFilePath + args[0].(string)
	if len(args) == 1 {
		Instance().Error(logstr)
	} else {
		Instance().Error(logstr, args[1:])
	}
}

func Errorf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	// 打印调用栈
	tfilePath := ""
	for i := 5; i > 0; i-- {
		_, tfp, tno, ok := runtime.Caller(i)
		if !ok {
			continue
		}
		tfilePath += fmt.Sprintf("%s:%d -> ", path.Base(tfp), tno)
	}
	logstr := tfilePath + format
	Instance().Errorf(logstr, args...)
}

func Fatal(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(args...)
		}
	}()
	Instance().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	Instance().Fatalf(format, args...)
}

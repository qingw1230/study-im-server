package log

import (
	"bufio"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var globalLogger *Logger

type Logger struct {
	*logrus.Logger
	Pid int
}

func init() {
	globalLogger = loggerInit("")
}

func NewPrivateLog(moduleName string) {
	globalLogger = loggerInit(moduleName)
}

func loggerInit(moduleName string) *Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.Level(config.Config.Log.RemainLogLevel))
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}
	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath"},
	})

	logger.AddHook(newFileLineHook())
	hook := NewLfsHook(time.Duration(config.Config.Log.RotationTime)*time.Hour, config.Config.Log.RemainRotationCount, moduleName)
	logger.AddHook(hook)
	return &Logger{Logger: logger, Pid: os.Getpid()}
}

func NewLfsHook(rotationTime time.Duration, maxRemainNum uint, moduleName string) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "debug", moduleName),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "info", moduleName),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "warn", moduleName),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "error", moduleName),
	}, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath"},
	})
	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string, moduleName string) *rotatelogs.RotateLogs {
	if moduleName != "" {
		moduleName = moduleName + "."
	}
	writer, err := rotatelogs.New(
		config.Config.Log.StorageLocation+moduleName+level+"."+"%Y-%m-%d",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
	} else {
		return writer
	}
}

func Info(args ...interface{}) {
	globalLogger.WithFields(logrus.Fields{
		"PID": globalLogger.Pid,
	}).Infoln(args...)
}

func Error(args ...interface{}) {
	globalLogger.WithFields(logrus.Fields{
		"PID": globalLogger.Pid,
	}).Errorln(args...)
}

func Debug(args ...interface{}) {
	globalLogger.WithFields(logrus.Fields{
		"PID": globalLogger.Pid,
	}).Debugln(args...)
}

func Warn(args ...interface{}) {
	globalLogger.WithFields(logrus.Fields{
		"PID": globalLogger.Pid,
	}).Warnln(args...)
}

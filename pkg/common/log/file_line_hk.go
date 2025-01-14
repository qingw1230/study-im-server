package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type fileLineHook struct {
}

func newFileLineHook() *fileLineHook {
	return &fileLineHook{}
}

func (f *fileLineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f *fileLineHook) Fire(entry *logrus.Entry) error {
	entry.Data["FilePath"] = findCaller(6)
	return nil
}

func findCaller(skip int) string {
	file, line := "", 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "log") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}

	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

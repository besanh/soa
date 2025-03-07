package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/gookit/slog"
	"github.com/gookit/slog/rotatefile"
)

func Errorf(format string, err ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	slog.WithFields(slog.M{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Errorf(format, err...)
}

func Debugf(format string, value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	slog.WithFields(slog.M{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Debugf(format, value...)
}

func InitLogger(level string, logFile string) {
	logLevel := slog.DebugLevel
	switch level {
	case "debug":
		logLevel = slog.DebugLevel
	case "error":
		logLevel = slog.ErrorLevel
	}
	slog.SetLogLevel(logLevel)
	logTemplate := "[{{level}}] Message: {{message}} {{data}} \n"

	slog.SetFormatter(slog.NewTextFormatter(logTemplate).WithEnableColor(true))
	writer, err := rotatefile.NewConfig(logFile).Create()
	if err != nil {
		panic(err)
	}

	log.SetOutput(writer)
}

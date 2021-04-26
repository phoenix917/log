package log

import (
	"github.com/larspensjo/config"
	"github.com/lestrrat/go-file-rotatelogs"
	fm "github.com/phoenix917/log/formatter"
	hk "github.com/phoenix917/log/hooks"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var Logger *logrus.Logger

func init() {
	fileName, level := customConfig()
	logPath := path.Join(currentDirectory(), "log")
	logPattern := path.Join(logPath, fileName)

	_, err := os.Stat(logPath)
	if err != nil && os.IsNotExist(err) {
		_ = os.Mkdir(logPath, os.ModePerm)
	}
	Logger = logrus.New()
	Logger.Hooks.Add(hk.NewContextHook())
	Logger.SetFormatter(&fm.CustomFormatter{})

	writer, err := rotatelogs.New(
		logPattern,
		// WithRotationTime设置日志分割的时间，这里设置为一天分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(time.Hour*24*30),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	Logger.SetOutput(writer)
	Logger.SetLevel(level)
}

func currentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func customConfig() (fileName string, level logrus.Level) {
	cfg, err := config.ReadDefault("config.ini")

	if cfg == nil || err != nil {
		return "", logrus.InfoLevel
	}

	fn, _ := cfg.String("log", "filename")

	l, err := cfg.String("log", "level")

	if err != nil {
		return fn, logrus.InfoLevel
	}

	switch l {
	case "":
		return fn, logrus.InfoLevel
	case "info":
		return fn, logrus.InfoLevel
	case "debug":
		return fn, logrus.DebugLevel
	case "error":
		return fn, logrus.ErrorLevel
	}

	return fn, logrus.InfoLevel
}

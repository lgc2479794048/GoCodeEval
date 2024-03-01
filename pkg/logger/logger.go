package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)
	Log.SetReportCaller(true)

	// 创建日志文件夹
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	// 配置日志文件输出
	logFilePath := filepath.Join(logDir, "go_code_eval.log")
	Log.SetOutput(&lumberjack.Logger{
		Filename:   logFilePath, // 日志文件的位置
		MaxSize:    10,          // 日志文件的最大大小（MB）
		MaxBackups: 3,           // 最多保留3个备份
		MaxAge:     28,          // 备份的最大天数
		Compress:   true,        // 是否压缩/归档旧文件
	})

	// 你也可以添加一个钩子输出到控制台
	Log.AddHook(&writerHook{
		Writer:    os.Stdout,
		LogLevels: logrus.AllLevels,
	})
}

// writerHook 是一个logrus钩子，用于同时将日志输出到指定的io.Writer
type writerHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Info 打印信息级别的日志
func Info(message string) {
	Log.Info(message)
}

// Warn 打印警告级别的日志
func Warn(message string) {
	Log.Warn(message)
}

// Error 打印错误级别的日志
func Error(message string) {
	Log.Error(message)
}

// Debug 打印调试级别的日志
func Debug(message string) {
	Log.Debug(message)
}

// Fatal 打印致命错误级别的日志，并退出程序
func Fatal(message string) {
	Log.Fatal(message)
}

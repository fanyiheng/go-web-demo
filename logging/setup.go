package logging

import (
	"fmt"
	"github.com/fanyiheng/go-web-demo/setting"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func Setup() {
	var (
		outputFile = path.Join("log","app")
		level      = logrus.InfoLevel.String()
		stdout     = false
	)
	if setting.LogSetting == nil {
		fmt.Println("conf.LogSetting is nil, use default logging config")
	} else {
		if setting.LogSetting.OutputFile == "" {
			fmt.Println("logging output file is not configured, use default")
		} else {
			outputFile = setting.LogSetting.OutputFile
		}
		if setting.LogSetting.Level == "" {
			fmt.Println("logging level is not configured, use default")
		} else {
			level = setting.LogSetting.Level
		}
		stdout = setting.LogSetting.Stdout
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("invalid logging level %s, use default level %s", setting.LogSetting.Level, logrus.InfoLevel.String())
		logLevel = logrus.InfoLevel
	}
	fmt.Printf("log used config: outputFile=%s, level=%s, stdout=%t\n", outputFile, logLevel.String(), stdout)
	logrus.SetLevel(logLevel)

	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		outputFile+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(outputFile+".log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// 如果日志不输出到控制台，则将默认日志输出到文件
	if !stdout {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:03:04.000",
		})
		logrus.SetOutput(logWriter)
	} else {
		lfHook := lfshook.NewHook(logWriter, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:03:04.000",
		})
		logrus.AddHook(lfHook)
	}
}

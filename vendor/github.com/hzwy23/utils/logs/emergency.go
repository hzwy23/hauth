package logs

import (
	"fmt"

	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var (
	back_emc = new(zap.SugaredLogger)
)

func Infow(msg string, keysAndValues ...interface{}) {
	back_emc.Infow(msg, keysAndValues...)
}

func init() {
	//GetDetails log dir from environment
	cfg := zap.NewProductionConfig()

	// 生成日志文件路径
	logpath := os.Getenv("HBIGDATA_HOME")
	filename := filepath.Join(logpath, "temp", "log", "emergency.log")

	// 判断日志所在目录,是否存在
	_, err := os.Stat(filepath.Join(logpath, "temp", "log"))
	if err != nil {
		if os.IsNotExist(err) {
			// 创建日志目录
			err := os.MkdirAll(filepath.Join(logpath, "temp", "log"), os.ModeDir)
			if err != nil {
				fmt.Errorf("%s", "文件不存在,创建日志文件失败")
				// 日志文件无法创建
				// 使用console作为日志输出
			} else {
				cfg.OutputPaths = []string{filename}
				cfg.ErrorOutputPaths = []string{filename}
			}
		}
		// 如果日志文件存在,但是无法获取Stat信息
		// 将日志输出到console上
	} else {
		cfg.OutputPaths = []string{filename}
		cfg.ErrorOutputPaths = []string{filename}
	}

	cfg.EncoderConfig.EncodeTime = iso8601TimeEncoder
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.Level.UnmarshalText([]byte("debug"))
	logger, err := cfg.Build()
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	back_emc = logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

package logs

import (
	"fmt"

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
	cfg.OutputPaths = []string{"./temp/log/emergency.log"}
	cfg.ErrorOutputPaths = []string{"./temp/log/emergency.log"}
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

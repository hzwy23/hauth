// Copyright 2017 hzwy23 Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"fmt"
	"time"

	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/hzwy23/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  = new(zap.SugaredLogger)
	lock = new(sync.RWMutex)
)

// 没有参数时,返回系统正常操作日志接口
// 当有参数时,不管参数是什么值,返回紧急日志备份接口
func GetLogger() *zap.SugaredLogger {
	lock.RLock()
	defer lock.RLock()
	return log
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	log.Error(v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	log.Warn(v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	log.Info(v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	log.Debug(v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

func iso8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func init() {
	//GetDetails log dir from environment
	logpath := os.Getenv("HBIGDATA_HOME")

	conf, err := config.GetConfig(path.Join(logpath, "conf", "asofdate.conf"))
	if err != nil {
		fmt.Errorf("%v", err)
		log = back_emc
		return
	}

	file, err := conf.Get("Hauth.log.file")
	if err != nil {
		fmt.Errorf("%s", "cant not find Hauth.log.file. so set default file.")
		file = "hauth.log"
	}

	filename := filepath.Join(logpath, "temp", "log", file)
	log_level, err := conf.Get("Hauth.log.level")
	if err != nil {
		fmt.Errorf("%s", "log level not set, set log level as default value.")
		log_level = "info"
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{filename}
	cfg.ErrorOutputPaths = []string{filename}

	cfg.EncoderConfig.EncodeTime = iso8601TimeEncoder
	cfg.DisableStacktrace = true
	cfg.Level.UnmarshalText([]byte(log_level))
	logger, err := cfg.Build()
	if err != nil {
		fmt.Errorf("%v", err)
		log = back_emc
		return
	}

	lock.Lock()
	defer lock.Unlock()
	log = logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

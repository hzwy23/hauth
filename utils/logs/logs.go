// Copyright 2014 beego Author. All Rights Reserved.
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
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/asofdate/utils/config"
)

// Log levels to control the logging output.
//const (
//	LevelEmergency = iota
//	LevelAlert
//	LevelCritical
//	LevelError
//	LevelWarning
//	LevelNotice
//	LevelInformational
//	LevelDebug
//)

// BeeLogger references the used application logger.
var Lg = logs.NewLogger()

func GetBeeLogger() *logs.BeeLogger {
	return Lg
}

// Emergency logs a message at emergency level.
func Emergency(v ...interface{}) {
	Lg.Emergency(generateFmtStr(len(v)), v...)
}

// Alert logs a message at alert level.
func Alert(v ...interface{}) {
	Lg.Alert(generateFmtStr(len(v)), v...)
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	Lg.Critical(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	Lg.Error(generateFmtStr(len(v)), v...)
}

// Warning logs a message at warning level.
func Warning(v ...interface{}) {
	Lg.Warning(generateFmtStr(len(v)), v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	Lg.Warn(generateFmtStr(len(v)), v...)
}

// Notice logs a message at notice level.
func Notice(v ...interface{}) {
	Lg.Notice(generateFmtStr(len(v)), v...)
}

// Informational logs a message at info level.
func Informational(v ...interface{}) {
	Lg.Informational(generateFmtStr(len(v)), v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	Lg.Info(generateFmtStr(len(v)), v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	Lg.Debug(generateFmtStr(len(v)), v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(v ...interface{}) {
	Lg.Trace(generateFmtStr(len(v)), v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Fatal(v ...interface{}) {
	Lg.Critical(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}

func init() {

	//GetDetails log dir from environment
	logpath := os.Getenv("HBIGDATA_HOME")

	conf, err := config.GetConfig(path.Join(logpath, "conf", "system.properties"))
	if err != nil {
		fmt.Println(err, "get system.properties failed.")
		Lg.Error(err.Error())
		return
	}

	file, err := conf.Get("Hauth.log.file")
	if err != nil {
		Lg.Warn("cant not find Hauth.log.file. so set default file.")
		file = "platform.log"
	}

	//log file name is platform.log
	filename := path.Join(logpath, "log", file)
	smaxlines, _ := conf.Get("Hauth.log.maxlines")
	smaxsize, _ := conf.Get("Hauth.log.maxsize")
	daily, _ := conf.Get("Hauth.log.daily")
	smaxdays, _ := conf.Get("Hauth.log.maxdays")
	filename = strings.Replace(filename, "\\", "/", -1)
	// set log config

	err = Lg.SetLogger("file", `{"filename":"`+filename+`","maxlines":`+smaxlines+`,"maxsize":`+smaxsize+`,"daily":`+daily+`,"maxdays":`+smaxdays+`}`)

	Lg.EnableFuncCallDepth(true)

	Lg.SetLogFuncCallDepth(3)

	lvl, err := conf.Get("Hauth.log.level")

	if err != nil {
		logs.Warn("get loglevel failed. set loglevel equal 3", err)
		Lg.SetLevel(3)
		Debug("init log, platform log dir is :", filename)
		return
	}

	LOGLVL, _ := strconv.Atoi(lvl)

	Lg.SetLevel(LOGLVL)

	Debug("init log, platform log dir is :", filename)
}

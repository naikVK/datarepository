//@author  vivek naik
//@version Thu Jul 05 2018 06:40:54 GMT+0530 (IST)

// Package loggermdl will help you log error
package loggermdl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

// Init  Init Logger
// maxBackupFileSize,  megabytes
// maxAgeForBackupFile,  days
func Init(fileName string, maxBackupCnt, maxBackupFileSize, maxAgeForBackupFileInDays int, loglevel zapcore.Level) {
	os.MkdirAll(filepath.Dir(fileName), os.ModePerm)

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxBackupFileSize, // megabytes
		MaxBackups: maxBackupCnt,
		MaxAge:     maxAgeForBackupFileInDays, // days
	})

	// zap.AddStacktrace(
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "disableCaller":false,
	  "encoderConfig": {
	    "messageKey": "m",
	    "callerKey": "c",
	    "callerEncode": 0,
	    "timeKey": "t",
		"levelKey": "l",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	json.Unmarshal(rawJSON, &cfg)
	core := zapcore.NewCore(
		//enc, //
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		loglevel,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	defer logger.Sync()
	sugar = logger.Sugar()

}

// GetCallers will return callers chain
func GetCallers(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// +build !prod

//@author  vivek naik
//@version Thu Jul 05 2018 06:40:34 GMT+0530 (IST)

// Package loggermdl
package loggermdl

import (
	"fmt"
	"os"

	logging "github.com/op/go-logging"
	goon "github.com/shurcooL/go-goon"
)

var log = logging.MustGetLogger("mkcllogger")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfile} %{callpath:5} ▶ %{level:.4s} %{id:03x}%{color:reset}`,
)

func init() {

	log.ExtraCalldepth = 1
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)

	logging.SetBackend(backendLeveled, backendFormatter)

}

// LogDebug logs a message at level Debug on the standard logger.
func LogDebug(args ...interface{}) {
	log.Debug("", args)
}

// LogInfo logs a message at level Info on the standard logger.
func LogInfo(args ...interface{}) {
	log.Info("", args)
}

// LogWarn logs a message at level Warn on the standard logger.
func LogWarn(args ...interface{}) {
	log.Warning("", args)
}

// LogError logs a message at level Error on the standard logger.
func LogError(args ...interface{}) {
	log.Error("", args)
}

// LogPanic logs a message at level Panic on the standard logger.
func LogPanic(args ...interface{}) {
	log.Panic(args)
}

// // LogJSONObject Format string
// func LogJSONObject(pobj interface{}) {
// 	jsonByte, _ := json.Marshal(pobj)
// 	var objnew map[string]interface{}
// 	json.Unmarshal(jsonByte, &objnew)

// 	f := colorjson.NewFormatter()
// 	f.Indent = 2

// 	s, _ := f.Marshal(objnew)
// 	fmt.Println(string(s))
// }

// // LogJSONByte Format string
// func LogJSONByte(pobj []byte) {
// 	var objnew map[string]interface{}
// 	json.Unmarshal(pobj, &objnew)

// 	f := colorjson.NewFormatter()
// 	f.Indent = 2

// 	s, _ := f.Marshal(objnew)
// 	fmt.Println(string(s))
// }

// // LogJSONString Format string
// func LogJSONString(str string) {
// 	var objnew map[string]interface{}
// 	json.Unmarshal([]byte(str), &objnew)

// 	f := colorjson.NewFormatter()
// 	f.Indent = 2

// 	s, _ := f.Marshal(objnew)
// 	fmt.Println(string(s))
// }

// LogHRStart can end line with <<<
func LogHRStart() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
}

// LogHREnd can end line with <<<
func LogHREnd() {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<")
}

// LogSpot will print block
func LogSpot(args ...interface{}) {
	fmt.Println(">>>>>>>>>>STRT>>>>>>>>>>>>>")
	log.Info("", args)
	fmt.Println("<<<<<<<<<END<<<<<<<<<<<<<<")
}

// LogVars Prints variables with formatting
func LogVars(xvars ...interface{}) {
	for _, i := range xvars {
		goon.Dump(i)
	}
}

// TODO:This Function bring back later// LogTable will print data in table form
// func LogTable(data []interface{}) {
// 	t := gotabulate.Create(data)
// 	fmt.Println(t.Render("grid"))
// }

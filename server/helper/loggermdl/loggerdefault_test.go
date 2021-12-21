//@author  vivek naik
//@version Thu Jul 05 2018 06:40:10 GMT+0530 (IST)

package loggermdl

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func ExampleLogVars() {
	type Example struct {
		name string
	}
	a := Example{"name"}

	LogVars(a)
	//Output:
	//(loggermdl.Example)(loggermdl.Example{
	//	name: (string)("name"),
	//})
}

func ExampleLogError() {
	LogError("a")
	//Output:
	//
}
func ExampleLogInfo() {
	LogInfo("a")
	//Output:
	//
}
func ExampleLogWarn() {
	LogWarn("a")
	//Output:
	//
}
func ExampleLogDebug() {
	LogDebug("a")
	//Output:
	//
}
func TestLogJSONString(t *testing.T) {
	a := `{"Age": 1,"Name": "sometext"}`

	assert.NotPanics(t, func() { LogJSONString(a) }, "The code did  panic")

	//Output2:
	// {
	//   "Age": 1,
	//   "Name": "sometext"
	// }
}
func ExampleLogHREnd() {
	LogHREnd()
	//Output:
	//<<<<<<<<<<<<<<<<<<<<<<<
}

func TestLogJSONObject(t *testing.T) {
	type MySt struct {
		Name string
		Age  int
	}
	m := MySt{"sometext", 1}
	// jsonUnmarshal(json.Marshal(m), &m)

	assert.NotPanics(t, func() { LogJSONObject(m) }, "The code did  panic")
	//Output1:
	//{
	//   "Age": 1,
	//   "Name": "sometext"
	//}
}

func TestLogJSONByte(t *testing.T) { // func ExampleLogJSONByte() {

	type MySt struct {
		Name string
		Age  int
	}
	m := MySt{"sometext", 1}
	a, _ := json.Marshal(m)

	assert.NotPanics(t, func() { LogJSONByte(a) }, "The code did  panic")
	//Output1:
	// {
	//   "Age": 1,
	//   "Name": "sometext"
	// }

}
func ExampleLogHRStart() {
	LogHRStart()
	//Output:
	//>>>>>>>>>>>>>>>>>>>>>>>
}

func ExampleNegetiveGetCaller() {
	a := GetCallers(-323232323)
	fmt.Println(len(a))
	//output:
	//8
}

func ExampleInit() {
	Init("filename", 3, 7, 5, zapcore.DebugLevel)
	//output:
	//
}

func TestLogPanic(t *testing.T) {
	assert.Panics(t, func() { LogPanic("a") }, "The code did not panic")
	assert.NotPanics(t, func() { LogSpot("A") }, "The code did not panic")

}

// func TestLogTable(t *testing.T) {

// 	type MySt struct {
// 		Name string
// 		Age  int
// 	}
// 	m1 := MySt{"sometextm1", 1}
// 	m2 := MySt{"sometext m2", 13}
// 	ary := make([]interface{}, 2)
// 	ary = append(ary, m1)
// 	ary = append(ary, m2)

// 	assert.NotPanics(t, func() { LogTable(ary) }, "The code did  panic")

// }

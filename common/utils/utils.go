package utils

import (
	"reflect"
	"runtime"
	"strings"
)

type LogFormatter func(funcName string, messages ...string) string

func NewLogFormatter(prefix string) LogFormatter {
	return func(funcName string, messages ...string) string {
		var msg string
		if len(messages) > 0 {
			msg = strings.Join(messages, " ")
		}
		return "[" + prefix + "." + funcName + "] " + msg
	}
}

// GetFN function to read function name
func GetFN(i interface{}) string {
	splitStr := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strings.Replace(splitStr[len(splitStr)-1], "-fm", "", 1)
}

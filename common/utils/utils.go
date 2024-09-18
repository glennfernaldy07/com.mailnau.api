package utils

import (
	"encoding/base64"
	"encoding/json"
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

func DecodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}

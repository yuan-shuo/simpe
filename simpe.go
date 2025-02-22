package simpe

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// track error link
func De(err error) error {
	if err == nil {
		return nil
	}
	pc, fullFilePath, line, ok := runtime.Caller(1) // 获取调用De方法的文件名和行号
	if !ok {
		fullFilePath = "???"
		line = 0
	}
	function := runtime.FuncForPC(pc).Name()
	fileName := path.Base(fullFilePath)     // 获取文件名
	packageName := getPackageName(function) // 获取包名

	location := fmt.Sprintf("%s/%s:%d", packageName, fileName, line)
	return errors.Wrap(err, location+" -> ")
}

func ShowErr(err error) {
	if err != nil {
		log.Printf("[Error source]: %v", errors.Cause(err))
		log.Printf("[Error link]: %v", err)
	}
}

// getPackageName extracts the package name from a fully qualified function name.
func getPackageName(fullFuncName string) string {
	lastSlash := strings.LastIndex(fullFuncName, "/")
	if lastSlash == -1 {
		return ""
	}
	lastDot := strings.LastIndex(fullFuncName[lastSlash:], ".")
	if lastDot == -1 {
		return ""
	}
	return fullFuncName[lastSlash+1 : lastSlash+lastDot]
}

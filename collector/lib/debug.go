// Copyright 2020-03-02 Nubes-Bridge. All rights reserved.

package lib

import (
	"fmt"
	"runtime"
)

var IsLogInfo bool = false
var IsLogWarn bool = true

func LogWarn(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)
	if IsLogWarn {
		fmt.Printf("%s:%d %v", file, line, info)
	}
}

func LogInfoln(a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintln(a...)
	if IsLogInfo {
		fmt.Printf("%s:%d %v", file, line, info)
	}
}

func LogInfo(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)
	if IsLogInfo {
		fmt.Printf("%s:%d %v", file, line, info)
	}
}

func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)
	fmt.Printf("%s:%d %v", file, line, info)
}

func Trace() {
	pc := make([]uintptr,10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}



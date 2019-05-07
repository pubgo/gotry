package gotry

import (
	"fmt"
	"runtime"
)

type TaskFn func(args ...interface{}) *_task_fn

const callDepth = 2

func funcCaller() string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller"
	}
	return fmt.Sprintf("%s:%d ", file, line)
}

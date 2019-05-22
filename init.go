package gotry

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var goPath = build.Default.GOPATH
var srcDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "src"), string(os.PathSeparator))
var modDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "pkg", "mod"), string(os.PathSeparator))

const callDepth = 2

func funcCaller() string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller"
	}

	return strings.TrimPrefix(strings.TrimPrefix(fmt.Sprintf("%s:%d ", file, line), srcDir), modDir)
}

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

func Retry(num int, fn func()) (err error) {
	_t := fibonacci()
	for i := 0; i < num; i++ {
		if err = Try(fn).KErr(); err == nil {
			return
		}
		time.Sleep(time.Second * time.Duration(_t()))
	}
	return
}

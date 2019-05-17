package gotry

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

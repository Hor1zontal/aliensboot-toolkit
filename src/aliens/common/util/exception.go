package util

import (
	"runtime"
	"fmt"
	"aliens/log"
)

func PrintStackDetail() {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	stackInfo := fmt.Sprintf("%s", buf[:n])
	log.Error("panic stack info %s", stackInfo)

}
package util

import (
	"runtime"
	"aliens/log"
)

func CatchStackDetail() {
	if err := recover(); err != nil {
		PrintStackDetail()
	}
}

func PrintStackDetail() {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	log.Errorf("%s", buf[:n])
}
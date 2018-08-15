/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/7/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package exception

import (
	"runtime"
	"aliens/log"
	"aliens/config"
)

func GameException(data interface{}) {
	panic(data)
}



func CatchStackDetail() {
	if err := recover(); err != nil {
		PrintStackDetail(err)
	}
}

func PrintStackDetail(err interface{}) {
	if config.LenStackBuf > 0 {
		buf := make([]byte, config.LenStackBuf)
		n := runtime.Stack(buf, false)
		log.Errorf("%v: %s",err, buf[:n])
	} else {
		log.Error("%v", err)
	}
}
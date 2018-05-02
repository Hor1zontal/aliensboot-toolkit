/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/27
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import "aliens/log"

func main() {
	a := []int{}

	for i:=0;i<10;i++ {
		a=append(a, i)
	}

	log.Debug(a)

	a=a[1:len(a)]

	log.Debug(a)
}

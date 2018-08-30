/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package template

import "strings"

type ServiceMessage struct {
	PackageName string
	modules     map[string]*Module
}

func (this *ServiceMessage) EnsureModule(name string) *Module {
	if this.modules == nil {
		this.modules = make(map[string]*Module)
	}
	module := this.modules[name]
	if module == nil {
		module = &Module{Name:name, UName:strFirstToUpper(name), Handlers:make(map[int]*ProtoHandler), Pushs:make(map[int]string)}
		this.modules[name] = module
	}
	return module
}

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func strFirstToUpper(str string) string {
	f := str[0:1]
	t := str[1:]

	return strings.ToUpper(f) + t
}

type Module struct {
	Name string
	UName string
	Handlers map[int]*ProtoHandler
	Pushs map[int]string
}

type ProtoHandler struct {
	//Name string
	Desc string
	ORequest string
	OResponse string
	OPush string
}

func (this *ProtoHandler) IsValid() bool {
	return this.ORequest != ""
	//return this.ORequest != "" && this.OResponse != ""
}
